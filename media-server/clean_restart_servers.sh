#!/bin/bash
# Restarts servers and clean Redis data volume
./stop_servers.sh
rm -rf ./room-api/redis-volume
./start_servers.sh