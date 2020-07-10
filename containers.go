package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
)

type Container struct {
	Changed time.Time
	Status  string
	Name    string
}

type Containers map[string]Container

// Add a container to the list of monitored containers.
func (c *Containers) Add(id string, new Container) {
	// Is container already in containers?
	existing, ok := (*c)[id]

	if !ok {
		(*c)[id] = new

		return
	}

	// This is an old state (we know a newer state). Ignore and
	// return.
	if existing.Changed.After(new.Changed) {
		return
	}

	// Remember the new state.
	(*c)[id] = new
}

// Healthy if all containers are healthy.
func (c Containers) Healthy() bool {
	for _, container := range c {
		if container.Status != types.Healthy && container.Status != types.NoHealthcheck {
			return false
		}
	}

	return true
}

// Unhealthy if one of the containers get unhealthy.
func (c Containers) Unhealthy() error {
	for _, container := range c {
		if container.Status == types.Unhealthy {
			return fmt.Errorf(
				"%w: %s",
				unhealthyError,
				strings.Join(c.UnhealtyContainers(), ", "),
			)
		}
	}

	return nil
}

// NonHealtyContainers returns a list of container names of containers that are not healthy (yet).
func (c Containers) NonHealtyContainers() []string {
	var nonHealthy []string

	for _, container := range c {
		if container.Status != types.Healthy && container.Status != types.NoHealthcheck {
			nonHealthy = append(nonHealthy, container.Name)
		}
	}

	return nonHealthy
}

// UnhealtyContainers returns a list of container names of containers that are not healthy (yet).
func (c Containers) UnhealtyContainers() []string {
	var unhealthy []string

	for _, container := range c {
		if container.Status == types.Unhealthy {
			unhealthy = append(unhealthy, container.Name)
		}
	}

	return unhealthy
}
