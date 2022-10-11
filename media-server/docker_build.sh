#!/bin/bash
# Build the containers
(cd ./room-api; docker build -t room-api .)
(cd ./chatroom-server; docker build -t chatroom-server .)