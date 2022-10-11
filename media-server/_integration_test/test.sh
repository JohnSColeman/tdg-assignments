#!/bin/bash

# Test to exceed capacity
(cd ..; ./clean_restart_servers.sh)
docker run --rm -i grafana/k6 run - <exceed_capacity_create_room_script.js