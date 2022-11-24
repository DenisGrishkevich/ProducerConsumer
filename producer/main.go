package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"producer/app"
	"sync"
	"syscall"
)

func main(){
	wg := sync.WaitGroup{}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	filePath := "app/ids.txt"

	wg.Add(1)
	producer := app.NewProducer(filePath)
	go producer.Run(&wg, ctx)
	fmt.Println("Producer was started")

	wg.Wait()
}