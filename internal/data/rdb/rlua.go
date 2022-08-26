package rdb

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"
)

var GETDEL_HASH string
var STOCK_HASH string
var LOCK_DEL_HASH string

var once sync.Once

func load(c context.Context, filePath string) string {
	lua, er := os.ReadFile(filePath)
	if er != nil {
		log.Fatal(er)
	}
	hash, _ := Get().ScriptLoad(c, string(lua)).Result()
	return hash
}
func LoadScript() {
	once.Do(func() {
		c := context.Background()
		LOCK_DEL_HASH = load(c, "internal/data/rdb/lua/lock_del.lua")
		STOCK_HASH = load(c, "internal/data/rdb/lua/stock.lua")
		GETDEL_HASH = load(c, "internal/data/rdb/lua/get_and_del.lua")
	})
}
func GetDel(ctx context.Context, key string) (interface{}, error) {
	r, er := Get().EvalSha(ctx, GETDEL_HASH, []string{key}).Result()
	return r, er
}

func StockDecr(ctx context.Context, key string, decr int) (int64, error) {
	r, er := Get().EvalSha(ctx, STOCK_HASH, []string{key}, []any{decr}).Result()
	number, _ := strconv.Atoi(r.(string))
	return int64(number), er
}

func LockDel(ctx context.Context, key string, val string) (interface{}, error) {
	return Get().EvalSha(ctx, LOCK_DEL_HASH, []string{key}, []any{val}).Result()
}
