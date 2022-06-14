package runner

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type DockerRunner struct{}

func (r *DockerRunner) ImagePull(containerImage string) error {
	// Pull the runner base image
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	out, err := cli.ImagePull(ctx, containerImage, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	defer out.Close()

	if _, err := io.Copy(os.Stdout, out); err != nil {
		return err
	}

	return nil
}

func (r *DockerRunner) RunContainer(containerName string, containerImage string, containerEnv []string, containerMounts []mount.Mount) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: containerImage,
		Env:   containerEnv,
		Tty:   false,
	}, &container.HostConfig{
		Mounts:      containerMounts,
		NetworkMode: "host",
	}, nil, nil, containerName)
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case status := <-statusCh:
		if status.StatusCode != 0 {
			return fmt.Errorf("container exited with status code: %d", status.StatusCode)
		}
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
	if err != nil {
		return err
	}

	if _, err := stdcopy.StdCopy(os.Stdout, os.Stderr, out); err != nil {
		return err
	}

	return nil
}

func (r *DockerRunner) GetContainerLogs(containerName string) (io.ReadCloser, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true}
	out, err := cli.ContainerLogs(ctx, containerName, options)
	if err != nil {
		return nil, err
	}

	return out, nil
}
