package rdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/ivpusic/grpool"
	"log"
	"strings"
	"sync"
	"time"
)

type RmqContext struct {
	context.Context
	Stream     string
	Group      string
	ConsumerId string
	Msg        redis.XMessage
	Ack        func()
}
type HandlerFunc func(rctx *RmqContext)

var handlerx = make(map[string]HandlerFunc)
var consumeOnce sync.Once

const salt = "#"
const prexfix = "MQ_"

func Send(stream string, msg map[string]interface{}) error {
	err := Get().XAdd(ctx, &redis.XAddArgs{
		Stream: fmt.Sprintf("%s%s", prexfix, stream),
		MaxLen: 10000,
		Approx: true,
		ID:     "*",
		Values: msg,
	}).Err()
	return err
}

func Handler(stream string, group string, handler HandlerFunc) {
	handlerx[stream+salt+group] = handler
}
func Run() {
	consumeOnce.Do(func() {
		for k, v := range handlerx {
			go receive(k, v)
		}
		log.Println("start consume...")
	})
}
func receive(key string, handler HandlerFunc) {
	addrx := strings.Split(key, salt)
	subject := fmt.Sprintf("%s%s", prexfix, addrx[0])
	group := addrx[1]

	consumer_id := uuid.NewString()
	Get().XGroupCreateMkStream(ctx, subject, group, "0")

	pool := grpool.NewPool(10, 5)
	defer pool.Release()

	go func() {
		for {
			rcli := Get()
			time.Sleep(time.Second * 1)
			pcmds, err := rcli.XPendingExt(ctx, &redis.XPendingExtArgs{
				Stream: subject,
				Group:  group,
				Idle:   time.Second * 60,
				Start:  "0",
				End:    "+",
				Count:  10,
				//Consumer: consumer_id,
			}).Result()
			if err != nil {
				time.Sleep(time.Second * 5)
				continue
			}

			xdel_ids := []string{}
			for _, cmd := range pcmds {
				if cmd.RetryCount > 6 {
					xdel_ids = append(xdel_ids, cmd.ID)
				} else {
					xmsgs, err := rcli.XRangeN(ctx, subject, cmd.ID, cmd.ID, 1).Result()
					if err != nil {
						log.Println(err)
					}
					if len(xmsgs) > 0 {
						rcli.XClaim(ctx, &redis.XClaimArgs{
							Stream:   subject,
							Group:    group,
							Consumer: cmd.Consumer,
							MinIdle:  0,
							Messages: []string{cmd.ID},
						})
						msg := xmsgs[0]
						pool.JobQueue <- func() {
							handler(&RmqContext{
								Context:    context.Background(),
								Stream:     subject,
								Group:      group,
								ConsumerId: consumer_id,
								Msg:        msg,
								Ack: func() {
									_, e := rcli.XAck(ctx, subject, group, msg.ID).Result()
									if e == nil {
										rcli.XDel(ctx, subject, msg.ID)
									}
								},
							})
						}
					}
				}
			}
			if len(xdel_ids) > 0 {
				rcli.XAck(ctx, subject, group, xdel_ids...)
				rcli.XDel(ctx, subject, xdel_ids...)
			}
			time.Sleep(time.Second * 5)
		}
	}()
	for {
		entries, err := Get().XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumer_id,
			Streams:  []string{subject, ">"},
			Count:    10,
			Block:    time.Second * 25,
			NoAck:    false,
		}).Result()
		if err != nil {
			time.Sleep(time.Second * 2)
		}
		if len(entries) == 0 {
			continue
		}
		msgs := entries[0].Messages
		for _, msg := range msgs {
			pool.JobQueue <- func() {
				defer func() {
					if err := recover(); err != nil {
						log.Println(fmt.Sprintf("subject:%v, err:%v", subject, err))
					}
				}()
				handler(&RmqContext{
					Context:    context.Background(),
					Stream:     subject,
					Group:      group,
					ConsumerId: consumer_id,
					Msg:        msg,
					Ack: func() {
						_, e := Get().XAck(ctx, subject, group, msg.ID).Result()
						if e == nil {
							Get().XDel(ctx, subject, msg.ID)
						}
					},
				})
			}
		}
	}
}
