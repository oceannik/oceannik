package connectors

import (
	"fmt"
	"log"

	"github.com/oceannik/oceannik/auth"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type AgentConnector struct {
	Conn *grpc.ClientConn
	// Ctx       context.Context
	// CtxCancel context.CancelFunc
}

func getServerAddr() string {
	serverHost := viper.GetString("client.agent_host")
	serverPort := viper.GetInt("client.agent_port")
	serverAddr := fmt.Sprintf("%s:%d", serverHost, serverPort)

	return serverAddr
}

func (ac *AgentConnector) Open() {
	caCertPath := viper.GetString("client.certs.ca_cert_path")
	certPath := viper.GetString("client.certs.cert_path")
	keyPath := viper.GetString("client.certs.key_path")
	isServer := false

	tlsCreds, err := auth.LoadTLSCreds(caCertPath, certPath, keyPath, isServer)
	if err != nil {
		log.Fatal("[Oceannik Client] Failed loading the TLS certs: ", err)
	}

	serverAddr := getServerAddr()
	log.Printf("[Ocean] Connecting to Agent at %s...", serverAddr)

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(tlsCreds), grpc.WithBlock())
	// conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("gRPC Dial failed: %v", err)
	}
	ac.Conn = conn
	// defer conn.Close()

	// client := pb.NewNamespaceServiceClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	// ac.Ctx = ctx
	// ac.CtxCancel = cancel

	// log.Printf("Default timeout: %ds", defaultTimeout)
	// defer cancel()
}

func (ac *AgentConnector) Close() {
	ac.Conn.Close()
	// defer ac.CtxCancel()
}

func (ac *AgentConnector) GetDeploymentServiceClient() pb.DeploymentServiceClient {
	client := pb.NewDeploymentServiceClient(ac.Conn)
	return client
}

func (ac *AgentConnector) GetNamespaceServiceClient() pb.NamespaceServiceClient {
	client := pb.NewNamespaceServiceClient(ac.Conn)
	return client
}

func (ac *AgentConnector) GetProjectServiceClient() pb.ProjectServiceClient {
	client := pb.NewProjectServiceClient(ac.Conn)
	return client
}

func (ac *AgentConnector) GetSecretServiceClient() pb.SecretServiceClient {
	client := pb.NewSecretServiceClient(ac.Conn)
	return client
}
