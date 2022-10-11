# Media Server
This project provides a "media server" in the form of a [web chatroom](./chatroom-server) hub and 
a [room API](./room-api) to provided access to rooms mediated via a redis node.

A use case would be for an embedded support feature where a customer needs to have an ad-hoc conversation with
numerous customer support agents to deal with an ongoing issue that requires immediate attention.

The development philosophy of KISS and YAGNI, code and software architecture are minimalistic.

*Ideally implement a proper media server such as [OvenMediaEngine](https://www.ovenmediaengine.com/ome), however this
has its own architecture and scaling solution that doesn't fit the problem and 1 week deadline. Instead, let's
adapt a simple web chatroom hub using some existing OSS, but note this is not scalable as discussed in Solution issues.

## Capacity requirements
According to the requirements:
- 30 media server instances
- max 300 users per server
- max 300 pax per room
- support for 10,000 concurrent users

## Fallback strategy
When there is no available capacity the server with the lowest message rate will be nominated for hosting the overflow
of new rooms.

## Prerequisites
Please install as required:
- Linux or WSL
- latest Go
- docker+docker-compose

## Solution issues
The systems architecture where rooms are bound to a single server instance rather than being virtualised creates
numerous inefficiencies and technical problems:
- without knowing the number of guests for a room in advance server over capacity can't be guaranteed
- given the above servers may fill to capacity leading to failure of guests to join a room
- if the above problem is resolved by attempting to plan room size that can result in under utilisation of servers

## Recommendations
Preferred solution would be able to scale in a manner where room membership is virtualised rather than bound to a
specific server.

## Scripts
Please check the prerequisites are satisfied and then build the containers.

build containers: `./docker_build.sh`

start servers: `./start_servers.sh`

stop servers: `./stop_servers.sh`

## Tests
[integration_load_tests](./_integration_test)

