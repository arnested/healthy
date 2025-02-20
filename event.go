package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	errTimeout   Error = "timeout while waiting for containers to be healthy"
	errUnhealthy Error = "containers are unhealthy"
)

func listen(containers Containers, since time.Time, timeout time.Duration, failOnUnhealthy bool) (bool, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return false, fmt.Errorf("creating Docker client: %w", err)
	}

	filter := filters.NewArgs()
	filter.Add("type", "container")
	filter.Add("event", "health_status")

	for id := range containers {
		filter.Add("container", id)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	msgs, errs := cli.Events(ctx, events.ListOptions{
		Filters: filter,
		Since:   strconv.FormatInt(since.Unix(), 10),
	})

	timeoutChan := time.After(timeout)

	for {
		select {
		case err := <-errs:
			return false, err
		case msg := <-msgs:
			container := Container{
				Status:  msg.Status[15:],
				Changed: time.Unix(msg.Time, msg.TimeNano),
			}

			containers.Add(msg.ID, container)

			if containers.Healthy() {
				return true, nil
			}

			if err := containers.Unhealthy(); err != nil && failOnUnhealthy {
				return false, err
			}
		case <-timeoutChan:
			return false, fmt.Errorf(
				"%w (%s): %s",
				errTimeout,
				timeout,
				strings.Join(containers.NonHealtyContainers(), ", "),
			)
		}
	}
}
