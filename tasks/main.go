package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeSetCounter = "testing:setcounter"
)

type SetCounterPayload struct {
	Counter int
}

var redisClient = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
})

var pool = goredis.NewPool(redisClient)
var rs = redsync.New(pool)

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewSetCounterTask(counter int) (*asynq.Task, error) {
	payload, err := json.Marshal(SetCounterPayload{Counter: counter})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSetCounter, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.
//---------------------------------------------------------------

func HandleSetCounterTask(ctx context.Context, t *asynq.Task) error {
	var p SetCounterPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	mutexname := fmt.Sprintf("counter-%d", p.Counter)
	mutex := rs.NewMutex(mutexname)
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.Lock(); err != nil {
		if err == redsync.ErrFailed {
			fmt.Println("Could not obtain lock!")
		} else {
			log.Println(err.Error())
		}
		return nil
	}
	log.Printf("Counter now: counter=%d", p.Counter)
	time.Sleep(time.Minute)
	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		log.Println("unlock failed")
	}
	return nil
}
