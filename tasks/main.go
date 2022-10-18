package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeSetCounter = "testing:setcounter"
)

type SetCounterPayload struct {
	Counter int
}

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
	log.Printf("Counter now: counter=%d", p.Counter)
	return nil
}
