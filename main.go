package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	since := time.Now()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [flags] [container_id_or_name ...]\n\nflags:\n", os.Args[0])
		flag.PrintDefaults()
	}

	failOnUnhealthy := flag.Bool("fail-on-unhealthy", false, "fail on unhealthy")
	timeout := flag.Duration("timeout", time.Hour, "timeout after waiting")

	flag.Parse()

	cli, err := client.NewEnvClient()
	if err != nil {
		fail(err)
	}

	c := Containers{}

	for _, container := range flag.Args() {
		ID, container, err := containerInfo(container, cli, since)

		if err != nil {
			fail(err)
		}

		c.Add(ID, *container)
	}

	if c.Healthy() {
		os.Exit(0)
	}

	if err := c.Unhealthy(); err != nil && *failOnUnhealthy {
		fail(err)
	}

	if _, err := listen(c, since, *timeout, *failOnUnhealthy); err != nil {
		fail(err)
	}
}

func containerInfo(containerID string, cli *client.Client, since time.Time) (string, *Container, error) {
	info, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "", nil, err
	}

	state := types.NoHealthcheck

	if info.State.Health != nil {
		state = info.State.Health.Status
	}

	c := &Container{
		Status:  state,
		Changed: since,
		Name:    strings.TrimLeft(info.Name, "/"),
	}

	return info.ID, c, nil
}

func fail(err error) {
	fmt.Fprintf(flag.CommandLine.Output(), "%s", err)
	os.Exit(1)
}
