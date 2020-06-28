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

If you specify the option `-fail-on-unhealthy` the command will exit with a
non-zero return code if one of the containers turn unhealthy, but will otherwise
wait for all containers to become healthy.
