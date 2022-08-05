package rdb

import (
	"context"
	"encoding/json"
	"github.com/eininst/fastgo/configs"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
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
var redisOnce sync.Once

var ctx = context.TODO()

func Get() *redis.Client {
	return rcli
}

func Setup() {
	var rconf redisConf
	rstr := configs.Get("redis").String()
	_ = json.Unmarshal([]byte(rstr), &rconf)
	redisOnce.Do(func() {
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
		log.Println("Connected to Redis server...")

		//LoadScript()
	})
}
