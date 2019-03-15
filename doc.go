/*

Command healthy waits for Docker container(s) to become healthy.

The command takes one or more container ID's as argument(s) and will
not exit until all of them are reported "healthy" as it Health check
status.

A container with no health check defined is always considered healthy.

When all specified containers are reported healthy the command will
exit with return code 0.

If you specify the option `-fail-on-unhealthy` the command will exit
with a non-zero return code if one of the containers turn unhealthy,
but will otherwise wait for all containers to become healthy.

*/
package main
