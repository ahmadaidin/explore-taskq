package queue

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

var QueueFactory = redisq.NewFactory()
var RedisClient = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
})

var MainQueue = QueueFactory.RegisterQueue(&taskq.QueueOptions{
	Name:  "queue_hitung_nilai",
	Redis: RedisClient,
})

func counter() error {
	nilai := 0
	ctx := context.Background()
	for i := 0; i < 10000000; {
		nilai += i
		RedisClient.Set(ctx, "nilai_sekarang", nilai, 3*time.Minute)
		i++
	}
	return nil
}

var CountTask = taskq.RegisterTask(&taskq.TaskOptions{
	Name:    "counter",
	Handler: counter,
})
