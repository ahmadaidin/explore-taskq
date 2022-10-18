package main

import (
	"context"
	"log"

	"github.com/ahmadaidin/explore-taskq/queue"
)

func main() {
	ctx := context.Background()

	if err := queue.MainQueue.Consumer().Start(ctx); err != nil {
		log.Fatal(err)
	}
}
