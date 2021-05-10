package runner

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/spf13/viper"
)

type DockerRunner struct{}

func getBaseImageName() string {
	baseImage := viper.GetString("agent.runner_base_image")

	return baseImage
}

func (r *DockerRunner) Prepare() {
	// Pull the runner base image
	// ctx := context.Background()
	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	// if err != nil {
	// 	panic(err)
	// }

	// runnerBaseImage := getBaseImageName()
	// out, err := cli.ImagePull(ctx, runnerBaseImage, types.ImagePullOptions{})
	// if err != nil {
	// 	panic(err)
	// }

	// defer out.Close()

	// io.Copy(os.Stdout, out)
}

func (r *DockerRunner) Run(containerName string) {
	// runnerBaseImage := getBaseImageName()
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	// reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: getBaseImageName(),
		// Cmd:   []string{"echo", "hello world"},
		Tty: false,
	}, nil, nil, nil, containerName)
	if err != nil {
		log.Fatal(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatal(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
	if err != nil {
		log.Fatal(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

func (r *DockerRunner) GetLogs(containerName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true}
	out, err := cli.ContainerLogs(ctx, containerName, options)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, out)
}
