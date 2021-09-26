package consumer

import (
	"OCG_final_project_BE/api/config"
	"OCG_final_project_BE/api/model"
	"OCG_final_project_BE/system/sendgrid"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"OCG_final_project_BE/system/rbmq"

	"gorm.io/gorm"
)

// Worker defines a worker
type Worker struct {
	wg      *sync.WaitGroup
	mailer  sendgrid.Mailer
	rmqChan *rbmq.Rabbit
	ctx     context.Context
}

// NewWorker creates new worker and gets incoming parammeters
func NewWorker(ctx context.Context, wg *sync.WaitGroup, mailer sendgrid.Mailer, rmqChan *rbmq.Rabbit) *Worker {
	return &Worker{
		ctx:     ctx,
		wg:      wg,
		mailer:  mailer,
		rmqChan: rmqChan,
	}
}

// Start starts worker to process message
//  Processing logic:
//    1. Wait for message
//    2. Send email with mailer (Sendgrid client)
//    3. Update database (thankyou_email_sent) to prevent duplicated emails
func (w *Worker) Start() {
	if w.mailer == nil {
		fmt.Println("cannot start worker since mailer is nil")
		return
	}
	msgs, err := w.rmqChan.Consume()
	rbmq.FailOnError(err, "Failed to register a consumer")

	for {
		select {
		case msg := <-msgs:
			var mailCt sendgrid.EmailContent
			//convert byte to struct
			err := json.Unmarshal(msg.Body, &mailCt)
			if err != nil {
				fmt.Println("Can't unMarshal body msgs")
				continue
			}
			// call iface mailer with  sendgrid content
			err = w.mailer.Send(&mailCt)
			if err != nil {
				fmt.Println("Can't send msgs in sendgrid")
				continue
			}
			//  declare prepare statement via Gorm pkg  sendto to handle func
			//  and then send it to handle func
			tx := config.Database.Session(&gorm.Session{PrepareStmt: true})

			err = model.UpdateOrderAfterSend(tx, mailCt.StartTime, mailCt.EndTime)
			if err != nil {
				fmt.Println("Can't query to update report_send to to true")
			}
		case <-w.ctx.Done():
			fmt.Println("Exiting worker")
			w.wg.Done()
			return
		}
	}
}
