package main

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

func amqp(i int, wg *sync.WaitGroup, ctx context.Context) {
	fmt.Println("starting connection: ", i)
	c, err := amqp091.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		panic("connection error")
	}

	fmt.Println("giving connection time to establish: ", i)
	time.Sleep(time.Second * 5)

	<-ctx.Done()

	c.Close()
	fmt.Println("done closing connection: ", i)

	wg.Done()
	fmt.Println("after wg.Done(): ", i)
}

const n = 16

func main() {
	var wg sync.WaitGroup
	wg.Add(n)

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < n; i++ {
		go amqp(i, &wg, ctx)
	}

	cancel()
	fmt.Println("cancel done")

	wg.Wait()
	fmt.Println("wait done")

	fmt.Println("sleeping")
	time.Sleep(time.Hour)
}
