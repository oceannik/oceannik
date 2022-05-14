package connectors

import (
	"fmt"
	"log"
	"time"

	"github.com/oceannik/oceannik/common/auth"
	pb "github.com/oceannik/oceannik/proto"
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

	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatalf("[Ocean] Could not connect to the Agent! Are you sure the TLS certificates are valid? Error: %v", err)
	}
	ac.Conn = conn
}

func (ac *AgentConnector) Close() {
	ac.Conn.Close()
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