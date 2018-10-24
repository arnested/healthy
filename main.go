package main /* import "arnested.dk/go/healthy" */

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

func main() {
	failOnUnhealthy := flag.Bool("fail-on-unhealthy", false, "fail on unhealthy")
	flag.Parse()

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	allHealthy := false

	for !allHealthy {
		allHealthy = true
		for _, container := range flag.Args() {
			healthy, err := containerHealthy(container, cli, failOnUnhealthy)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if !healthy {
				allHealthy = false
			}
		}

		if !allHealthy {
			time.Sleep(2000 * time.Millisecond)
		}
	}

	if allHealthy {
		os.Exit(0)
	}
}

func containerHealthy(containerID string, cli *client.Client, failOnUnhealthy *bool) (bool, error) {
	healthStatus, err := containerHealthStatus(containerID, cli)
	if err != nil {
		return false, err
	}

	if *failOnUnhealthy == true && healthStatus == "unhealthy" {
		return false, errors.Errorf("%s is unhealthy", containerID)
	}

	if healthStatus == "healthy" {
		return true, nil
	}

	return false, nil
}

func containerHealthStatus(containerID string, cli *client.Client) (string, error) {
	info, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "error", err
	}

	if info.State.Health == nil {
		return "healthy", nil
	}

	return info.State.Health.Status, nil
}
