#!/bin/bash
# Start the containers
docker-compose -f ./chatroom-server/docker-compose.yml up -d
docker-compose -f ./room-api/docker-compose.yml up -d