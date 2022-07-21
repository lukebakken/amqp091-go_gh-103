package main

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

var wg sync.WaitGroup

func amqp(ctx context.Context) {
	defer func() {
		wg.Done()
	}()
	c, err := amqp091.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		panic("connection error")
	}
	defer c.Close()

	<-ctx.Done()
}

const n = 16

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(n)
	for i := 0; i < n; i++ {
		go amqp(ctx)
	}

	cancel()
	wg.Wait()

	time.Sleep(time.Hour)
}
