package main

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	// "sync"
	"time"
)

// var wg sync.WaitGroup

func amqp(i int, ctx context.Context) {
	fmt.Println("starting connection: ", i)
	c, err := amqp091.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		panic("connection error")
	}

	<-ctx.Done()

	fmt.Println("closing connection: ", i)
	// wg.Done()
	c.Close()
}

const n = 16

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// wg.Add(n)
	for i := 0; i < n; i++ {
		go amqp(i, ctx)
	}

	fmt.Println("cancel start")
	cancel()
	fmt.Println("cancel done")
	// wg.Wait()

	fmt.Println("sleeping")
	time.Sleep(time.Hour)
}
