package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})

var ctx = context.Background()

func authoriseRoom(rId roomId, sId serverId) (result string, error error) {
	return rdb.Set(ctx, rId, sId, roomJoinPeriod).Result()
}

func deauthoriseRoom(rId roomId) (result int64, error error) {
	return rdb.Del(ctx, rId).Result()
}

func messageRate(sId string) (float64, error) {
	smr, rerr := rdb.Get(ctx, sId).Result()
	if rerr != nil {
		return 0, rerr
	}
	return strconv.ParseFloat(smr, 32)
}

func closedRooms() []string {
	var rIds []string
	for {
		rId, rerr := rdb.LPop(ctx, "closed").Result()
		if rerr != nil {
			break
		}
		rIds = append(rIds, rId)
	}
	return rIds
}
