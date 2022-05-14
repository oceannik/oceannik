package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/docker/docker/api/types/mount"
	"github.com/oceannik/oceannik/agent/database"
	"github.com/oceannik/oceannik/agent/runner"
	"github.com/oceannik/oceannik/common/auth"
	pb "github.com/oceannik/oceannik/common/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	"net/http"
	_ "net/http/pprof"
)

var (
	protocol         = "tcp"
	databaseInitData = false
)

func getContainerNameForDeploymentId(id uint) string {
	return fmt.Sprintf("oceannik_runner.deployment_%d", id)
}

func deploymentWorker(db *gorm.DB, requestedDeploymentIds <-chan uint) {
	runnerCertsPath := viper.GetString("agent.runner.certs_path")
	containerImage := viper.GetString("agent.runner.base_image")
	containerVolumes := []mount.Mount{
		{
			Type:     mount.TypeBind,
			Source:   runnerCertsPath,
			Target:   "/usr/oceannik/user-certs",
			ReadOnly: true,
		},
	}
	defaultDeploymentStrategy := "blue-green"

	for deploymentId := range requestedDeploymentIds {
		log.Printf("[Worker] Got the request to process Deployment with ID: %d", deploymentId)

		startedAt := time.Now()
		status := pb.Deployment_STARTED.String()
		database.UpdateDeploymentStatus(db, deploymentId, status, startedAt, time.Time{})

		deployment, result := database.GetDeploymentByID(db, deploymentId)
		if result.Error != nil {
			fmt.Fprintf(os.Stderr, "a deployment could not be fetched: %s", result.Error)
			break
		}

		project, result := database.GetProjectByID(db, deployment.Project.ID)
		if result.Error != nil {
			fmt.Fprintf(os.Stderr, "a project could not be fetched: %s", result.Error)
			break
		}

		containerName := getContainerNameForDeploymentId(deploymentId)
		containerEnv := []string{
			fmt.Sprintf("OCEANNIK_DEPLOYMENT_STRATEGY=%s", defaultDeploymentStrategy),
			fmt.Sprintf("OCEANNIK_PROJECT_ID=%d", project.ID),
			fmt.Sprintf("OCEANNIK_PROJECT_NAME=%s", project.Name),
			fmt.Sprintf("OCEANNIK_PROJECT_REPO=%s", project.RepositoryUrl),
			fmt.Sprintf("OCEANNIK_PROJECT_REPO_BRANCH=%s", project.RepositoryBranch),
			fmt.Sprintf("OCEANNIK_SERVICE_CONFIG_PATH=%s", project.ConfigPath),
		}

		runnerEngine := runner.DockerRunner{}
		pullErr := runnerEngine.ImagePull(containerImage)
		if pullErr != nil {
			fmt.Fprintf(os.Stderr, "the base image could not be pulled: %s", pullErr)
		}

		runErr := runnerEngine.RunContainer(containerName, containerImage, containerEnv, containerVolumes)
		if runErr != nil {
			fmt.Fprintf(os.Stderr, "a container created by the Runner Engine has failed: %s", runErr)
			status = pb.Deployment_EXITED_FAILURE.String()
		} else {
			status = pb.Deployment_EXITED_SUCCESS.String()
		}

		exitedAt := time.Now()
		database.UpdateDeploymentStatus(db, deploymentId, status, startedAt, exitedAt)
	}
}

func Start(serverPort int, databasePath string, devMode bool, devServerHost string, devServerPort int) {
	if devMode {
		go func() {
			log.Println("[Developer Mode] Starting HTTP Server for application profiling.")
			log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", devServerHost, devServerPort), nil))
		}()
	}

	caCertPath := viper.GetString("agent.certs.ca_cert_path")
	certPath := viper.GetString("agent.certs.cert_path")
	keyPath := viper.GetString("agent.certs.key_path")
	isServer := true

	tlsCreds, err := auth.LoadTLSCreds(caCertPath, certPath, keyPath, isServer)
	if err != nil {
		log.Fatal("[Oceannik Agent] Failed loading the TLS certs: ", err)
	}

	// Configure the database, populate with data if needed
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		databaseInitData = true
	}
	db := database.Open(databasePath)
	database.PerformAutoMigrate(db)
	if databaseInitData {
		database.InitData(db)
	}

	maxCapacity := viper.GetInt("agent.deployments_queue_max_capacity")
	// Create a worker for processing deployment requests
	requestedDeploymentIds := make(chan uint, maxCapacity)
	defer close(requestedDeploymentIds)
	go deploymentWorker(db, requestedDeploymentIds)

	// Configure the gRPC server and services
	listener, err := net.Listen(protocol, fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatal(err)
	}

	deploymentServiceServer := DeploymentServiceServer{db: db, runnerChan: requestedDeploymentIds}
	namespaceServiceServer := NamespaceServiceServer{db: db}
	projectServiceServer := ProjectServiceServer{db: db}
	secretServiceServer := SecretServiceServer{db: db}

	server := grpc.NewServer(
		grpc.Creds(tlsCreds),
	)

	pb.RegisterDeploymentServiceServer(server, &deploymentServiceServer)
	pb.RegisterNamespaceServiceServer(server, &namespaceServiceServer)
	pb.RegisterProjectServiceServer(server, &projectServiceServer)
	pb.RegisterSecretServiceServer(server, &secretServiceServer)

	reflection.Register(server)

	log.Printf("[Oceannik Agent] gRPC Server is listening on %s/%s...", listener.Addr().String(), protocol)
	err = server.Serve(listener)
	log.Fatal(err)
}
