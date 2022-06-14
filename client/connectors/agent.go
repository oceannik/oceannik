package connectors

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oceannik/oceannik/common/auth"
	pb "github.com/oceannik/oceannik/common/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type AgentConnector struct {
	Conn *grpc.ClientConn
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
	dialTimeout := viper.GetInt("client.dial_timeout")
	log.Printf("[Ocean] Connecting to Agent at %s (timeout: %d seconds)...", serverAddr, dialTimeout)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		serverAddr,
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("[Ocean] Could not connect to the Agent! Are you sure the TLS certificates are valid? Error: %v", err)
	}
	ac.Conn = conn
}

func (ac *AgentConnector) Close() {
	err := ac.Conn.Close()
	if err != nil {
		log.Fatalf("[Ocean] Connection to Agent could not be closed. Error: %v", err)
	}
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
