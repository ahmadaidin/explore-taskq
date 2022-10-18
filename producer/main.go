// In your application code, import the above package and use Client to put tasks on queues.

package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ahmadaidin/explore-taskq/tasks"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	// ------------------------------------------------------
	// Example 1: Enqueue task to be processed immediately.
	//            Use (*Client).Enqueue method.
	// ------------------------------------------------------
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	x := r1.Intn(100)
	task, err := tasks.NewSetCounterTask(x)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	info, err = client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s for the second time", info.ID, info.Queue)
}
