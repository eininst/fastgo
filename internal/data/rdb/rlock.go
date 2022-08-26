package rdb

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type RedisLock struct {
	Ctx context.Context
	Key string
	Val string
}

const RLOCK_PREFIX = "RLOCK"

func NewLock(ctx context.Context, lockName string) *RedisLock {
	key := fmt.Sprintf("%s_%s", RLOCK_PREFIX, lockName)
	val := fmt.Sprintf("%s_%s", lockName, uuid.NewString())
	return &RedisLock{
		Ctx: ctx,
		Key: key,
		Val: val,
	}
}
func CurrentTime() int64 {
	return time.Now().UnixNano() / 1000000
}
func (rlock *RedisLock) Acquire(timeout time.Duration) bool {
	endtime := CurrentTime() + timeout.Milliseconds()
	for {
		if CurrentTime() > endtime {
			return false
		}
		ok, er := Get().SetNX(rlock.Ctx, rlock.Key, rlock.Val, timeout*2).Result()
		if er != nil {
			return false
		}
		if ok {
			return true
		}
		time.Sleep(time.Millisecond * 5)
	}

}

func (rlock *RedisLock) Release() bool {
	r, err := LockDel(rlock.Ctx, rlock.Key, rlock.Val)
	//r, err := Get().EvalSha(rlock.Ctx, LOCK_DEL_HASH, []string{rlock.Key}, []any{rlock.Val}).Result()
	if err != nil {
		return false
	}
	if reply, ok := r.(int64); !ok {
		return false
	} else {
		return reply == 1
	}
}
