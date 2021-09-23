package seed

import (
	"OCG_final_project_BE/system/consumer"
	"OCG_final_project_BE/system/rbmq"
	"OCG_final_project_BE/system/scheduler"
	"OCG_final_project_BE/system/sendgrid"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
)

func CreateSeed() {
	// prepare params
	rmqChan := rbmq.NewChannelMQ("test")
	ApiKey := os.Getenv("SENDGRID_KEY")

	// prepare db
	db, err := sql.Open("mysql", "root:duyngo99@tcp(127.0.0.1:3306)/testsearch?charset=utf8mb4&parseTime=True&loc=Local") //TODO Change that!
	if err != nil {
		panic(err)
	}
	defer db.Close()

	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())

	// sql
	sched := scheduler.NewScheduler(ctx, rmqChan)
	go func() {
		sched.Start()
	}()

	mailer := sendgrid.NewSendgrid(ApiKey)
	consumer := consumer.NewWorker(ctx, wg, db, mailer, rmqChan) //worker

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

	wg.Add(1) // add 1 for worker only. don't need for scheduler
	// run worker (as a receiver of msgExchange channel first)
	go consumer.Start()

	// wait for the worker finishes its job
	wg.Wait()
}
