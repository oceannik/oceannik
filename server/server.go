package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/oceannik/oceannik/auth"
	"github.com/oceannik/oceannik/database"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/oceannik/oceannik/runner"
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

func deploymentWorker(db *gorm.DB, requestedDeploymentIds <-chan uint) {
	for deploymentId := range requestedDeploymentIds {
		// TODO: Refactor
		log.Printf("[Worker] Got the request to process Deployment with ID: %d", deploymentId)

		startedAt := time.Now()
		status := pb.Deployment_STARTED.String()

		containerName := fmt.Sprintf("oceannik_runner.deployment_%d", deploymentId)
		r := runner.DockerRunner{}
		r.Prepare()
		r.Run(containerName)

		exitedAt := time.Now()
		status = pb.Deployment_EXITED_SUCCESS.String()
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

	// Create a worker for processing deployment requests
	requestedDeploymentIds := make(chan uint, 100)
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
