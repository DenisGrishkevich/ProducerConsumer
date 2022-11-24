package main

import (
	"consumer/app"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main(){
	wg := sync.WaitGroup{}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	
	wg.Add(1)
	consumer := app.NewConsumer()
	go consumer.Run(&wg, ctx)

	fmt.Println("All components started.")
	fmt.Println("Waiting for interrupt...")

	wg.Wait()
	fmt.Println("Program gracefully shut down.")
}