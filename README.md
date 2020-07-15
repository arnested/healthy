[![Maintainability](https://api.codeclimate.com/v1/badges/e897778bd75491914adc/maintainability)](https://codeclimate.com/github/arnested/go-healthy/maintainability)
[![Docker image size](https://badgen.net/docker/size/arnested/healthy)](https://hub.docker.com/r/arnested/healthy)
[![Build Status](https://travis-ci.com/arnested/go-healthy.svg?branch=master)](https://travis-ci.com/arnested/go-healthy)
[![Release](https://img.shields.io/github/release/arnested/go-healthy.svg)](https://github.com/arnested/go-healthy/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/arnested.dk/go/healthy/)](https://goreportcard.com/report/arnested.dk/go/healthy)
[![CLA assistant](https://cla-assistant.io/readme/badge/arnested/go-healthy)](https://cla-assistant.io/arnested/go-healthy)
[![GoDoc](https://godoc.org/arnested.dk/go/healthy?status.svg)](https://pkg.go.dev/arnested.dk/go/healthy)

# healthy
--
Command healthy waits for Docker container(s) to become healthy.

The command takes one or more container ID's as argument(s) and will not exit
until all of them are reported "healthy" as it Health check status.

A container with no health check defined is always considered healthy.

When all specified containers are reported healthy the command will exit with
return code 0.

    usage: healthy [flags] [container_id_or_name ...]

    flags:
      -fail-on-unhealthy
    	fail on unhealthy
      -timeout duration
    	timeout after waiting (default 1h0m0s)

In a docker-compose setup you could wait for all services to be healthy by
running:

    healthy $(docker-compose ps -q)

Or just could wait for the database service to be healthy by running:

    healthy $(docker-compose ps -q database)

To wait no longer than one and half minute for containers to be healthy:

    healthy -timeout 1m30s $(docker-compose ps -q)

You can also run healthy using a Docker image:

    docker run --rm -v /var/run/docker.sock:/var/run/docker.sock:ro arnested/healthy $(docker-compose ps -q)
