package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	timeoutError   Error = "timeout while waiting for containers to be healthy"
	unhealthyError Error = "containers are unhealthy"
)

func listen(c Containers, since time.Time, timeout time.Duration, failOnUnhealthy bool) (bool, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return false, err
	}

	filter := filters.NewArgs()
	filter.Add("type", "container")
	filter.Add("event", "health_status")

	for id := range c {
		filter.Add("container", id)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	msgs, errs := cli.Events(ctx, types.EventsOptions{
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

			c.Add(msg.ID, container)

			if c.Healthy() {
				return true, nil
			}

			if err := c.Unhealthy(); err != nil && failOnUnhealthy {
				return false, err
			}
		case <-timeoutChan:
			return false, fmt.Errorf(
				"%w (%s): %s",
				timeoutError,
				timeout,
				strings.Join(c.NonHealtyContainers(), ", "),
			)
		}
	}
}
