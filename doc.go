/*

Command healthy waits for Docker container(s) to become healthy.

The command takes one or more container ID's as argument(s) and will
not exit until all of them are reported "healthy" as it Health check
status.

A container with no health check defined is always considered healthy.

When all specified containers are reported healthy the command will
exit with return code 0.

	usage: healthy [flags] [container_id_or_name ...]

	flags:
	  -fail-on-unhealthy
		fail on unhealthy
	  -timeout duration
		timeout after waiting (default 1h0m0s)

In a docker-compose setup you could wait for all services to be
healthy by running:

	healthy $(docker-compose ps -q)

Or just could wait for the database service to be healthy by running:

	healthy $(docker-compose ps -q database)

To wait no longer than one and half minute for containers to be healthy:

	healthy -timeout 1m30s $(docker-compose ps -q)

You can also run healthy using a Docker image:

	docker run --rm -v /var/run/docker.sock:/var/run/docker.sock:ro arnested/healthy $(docker-compose ps -q)

*/
package main
