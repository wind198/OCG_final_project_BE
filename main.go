package main

import (
	"OCG_final_project_BE/api/router"
	"OCG_final_project_BE/system/consumer"
	"OCG_final_project_BE/system/rbmq"
	"OCG_final_project_BE/system/scheduler"
	"OCG_final_project_BE/system/sendgrid"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	fmt.Println("Start")
	// barchart.Test_CreateChartProduct(&testing.T{})
	rmqChan := rbmq.NewChannelMQ("rmq")
	ApiKey := os.Getenv("SENDGRID_KEY")
	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())

	sched := scheduler.NewScheduler(ctx, rmqChan)
	go func() {
		sched.Start()
	}()

	mailer := sendgrid.NewSendgrid(ApiKey)
	consumer := consumer.NewWorker(ctx, wg, mailer, rmqChan) //worker

	// Notify causes package signal to relay incoming signals to c.
	// if send signal interrupt Notify will receive and send to into c channel
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c // waits for the termination signal such as ctrl+c
		fmt.Printf("Got %s signal. Exiting...\n", sig)
		sched.Stop() // stop scheduler at the end
		cancelFunc()
	}()
	// start router
	go router.HandleRequests()
	wg.Add(1) // add 1 for worker only. don't need for scheduler
	// run worker (as a receiver of msgExchange channel first)
	go consumer.Start()

	// wait for the worker finishes its job
	wg.Wait()
}
