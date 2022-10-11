#!/bin/bash
# Stop the containers
docker-compose -f ./room-api/docker-compose.yml down
docker-compose -f ./chatroom-server/docker-compose.yml down