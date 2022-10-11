# chatroom-server 
This project is a simple text chatroom web server that we will serve as the "media server" implementation. 

The original [Golang-realtime-chat-rooms](https://github.com/AnupKumarPanwar/Golang-realtime-chat-rooms) code is by
Anup Kumar Panwar. Anups code has been modified to validate the request room IDs via redis and a healthcheck endpoint
will be served for a root context GET operation, the hub will also monitor and publish a message rate.

*Note that binding rooms to a single server leads to unpredictable room capacity and inefficient use of capacity.

## Handling overload
There are 2 strategies to handle overload.
1. when all connections to a room are closed this is published for cleanup
1. each server monitors and publishes a message rate which can be used to nominate a fallback server

## Start the server
To directly start a server: `go run .`
 
browse to: http://localhost:8080/room/1

## Docker
build: `docker build -t chatroom-server .`

run: `docker run --name chatroom -p 8080:8080 chatroom-server`

### Docker compose
Use docker-compose to run 30 instances fixed to port 50001 to 50030 as required.

`docker-compose up`