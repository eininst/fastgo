package rdb

import (
	"context"
	"encoding/json"
	"fastgo/configs"
	"github.com/eininst/flog"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type redisConf struct {
	Addr         string `json:"addr"`
	Db           int    `json:"db"`
	PoolSize     int    `json:"poolSize"`
	MinIdleCount int    `json:"minIdleCount"`
	Password     string `json:"password"`
}

var rcli *redis.Client

var ctx = context.TODO()

func Get() *redis.Client {
	return rcli
}

func New() *redis.Client {
	var rconf redisConf
	rstr := configs.Get("redis").String()
	_ = json.Unmarshal([]byte(rstr), &rconf)

	rcli = redis.NewClient(&redis.Options{
		Addr:         rconf.Addr,
		Password:     rconf.Password,
		DB:           rconf.Db,
		DialTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     rconf.PoolSize,
		MinIdleConns: rconf.MinIdleCount,
		PoolTimeout:  30 * time.Second,
	})
	_, err := rcli.Ping(ctx).Result()

	if err != nil {
		log.Fatal("Unbale to connect to Redis", err)
	}
	flog.Info("Connected to Redis server...")

	return rcli
}
