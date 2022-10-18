package main

import (
	"log"

	"github.com/ahmadaidin/explore-taskq/queue"
)

func main() {
	if err := queue.MainQueue.Consumer().Stop(); err != nil {
		log.Fatal(err)
	}
}
