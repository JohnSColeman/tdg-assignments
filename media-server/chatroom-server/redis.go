package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"os"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "host.docker.internal:6379",
	Password: "",
	DB:       0,
})

var invalidRoomError = errors.New("invalid room ID")
var incorrectServerError = errors.New("incorrect server")

var serverId = os.Getenv("SID")

var ctx = context.Background()

// verify the room ID is allocated to this server
func verifyRoomId(roomId string) error {
	sid, err := rdb.Get(ctx, roomId).Result()
	if err != nil {
		if err == redis.Nil {
			return invalidRoomError
		}
		return err
	}
	if sid != serverId {
		return incorrectServerError
	}
	return nil
}

// notification that a room is not currently required
func closeRoom(roomId string) error {
	_, err := rdb.RPush(ctx, "closed", roomId).Result()
	return err
}

func updateMessageRate(count float32) error {
	_, err := rdb.Set(ctx, serverId, count, 0).Result()
	return err
}
