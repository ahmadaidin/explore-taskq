package main

import (
	"context"
	"time"

	"github.com/ahmadaidin/explore-taskq/queue"
)

func main() {
	var ctx = context.Background()
	q := queue.MainQueue
	for i := 0; i < 10; i++ {
		// Add the task without any args.
		msg := queue.CountTask.WithArgs(ctx)
		msg.Delay = time.Second
		err := q.Add(msg)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
	q.Close()
}
