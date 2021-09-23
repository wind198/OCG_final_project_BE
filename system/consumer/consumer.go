package consumer

import (
	"OCG_final_project_BE/system/sendgrid"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	"OCG_final_project_BE/system/rbmq"
)

// Worker defines a worker
type Worker struct {
	wg      *sync.WaitGroup
	mailer  sendgrid.Mailer
	rmqChan *rbmq.Rabbit
	ctx     context.Context
	db      *sql.DB
}

// NewWorker creates new worker and gets incoming parammeters
func NewWorker(ctx context.Context, wg *sync.WaitGroup, db *sql.DB, mailer sendgrid.Mailer, rmqChan *rbmq.Rabbit) *Worker {
	return &Worker{
		ctx:     ctx,
		wg:      wg,
		mailer:  mailer,
		rmqChan: rmqChan,
		db:      db,
	}
}

// Start starts worker to process message
//  Processing logic:
//    1. Wait for message
//    2. Send email with mailer (Sendgrid client)
//    3. Update database (thankyou_email_sent) to prevent duplicated emails
func (w *Worker) Start() {
	if w.mailer == nil || w.db == nil {
		fmt.Println("cannot start worker since mailer is nil")
		return
	}
	sttm, err := w.db.Prepare("UPDATE `order` SET thankyou_email_sent = ? WHERE id =?;")
	if err != nil {
		fmt.Println("Can't prepare statement to update thankyou_email_sent")
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
			// excute query
			_, err = sttm.Query(true, mailCt.ID)
			if err != nil {
				fmt.Println("Can't query to update thankyou_email_sent to true")
			}
		case <-w.ctx.Done():
			fmt.Println("Exiting worker")
			w.wg.Done()
			return
		}
	}
}
