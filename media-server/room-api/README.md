# room-api
The room api acts as a governor to perform some basic load balancing of the media servers.

## Docker
build: `docker build -t room-api .`

run: `docker run --name rooms -p 3000:3000 room-api`

### Docker compose
Use docker-compose to run 30 instances fixed to port 50001 to 50030 as required.

`docker-compose up`