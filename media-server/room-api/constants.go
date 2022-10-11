package main

import "time"

// todo make configurable for production
const mediaServerFirstPort = 50001      // the first media server port
const mediaServerInstances = 30         // the number of media servers
const roomJoinPeriod = time.Minute * 30 // the duration a room can be joined after creation
