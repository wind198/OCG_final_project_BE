package scheduler

import (
	"OCG_final_project_BE/api/model"
	"OCG_final_project_BE/system/barchart"
	"OCG_final_project_BE/system/rbmq"
	"OCG_final_project_BE/system/sendgrid"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"time"

	maller "github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/robfig/cron/v3"
)

/*
Goal:
 - is to scan order table for getting new order to send thank you email
 - run at every minute
 - after send the email -> set thankyou_email_sent to true

Definition:
 - if created_at < now - 1 minutes && thankyou_email_sent == false -> schedule to send email
*/

const (
	DefaultReportSubject   = "Report for shop owner from Decorshop"
	DefaultReportBodyPlain = "Report for shop owner from Decorshop. Here's your reports today:"
	DefaultReportBodyHtml  = "<strong>Report for shop owner from Decorshop. Here's your reports today:</strong>"
	DefaultFromName        = "My Decor Shop"
	DefaultFromEmail       = "duynvwork@gmail.com"
	DefaultShopOwnerEmail  = "ngoduy9911@gmail.com"
	DefaultShopOwnerNamme  = "Duynv Decor"
)

type Scheduler struct {
	db   *sql.DB
	cron *cron.Cron
	// outChan chan<- *mail.EmailContent
	rmqChan *rbmq.Rabbit
	ctx     context.Context
}

func NewScheduler(ctx context.Context, db *sql.DB, rmqChan *rbmq.Rabbit) *Scheduler {
	return &Scheduler{
		ctx:     ctx,
		db:      db,
		cron:    cron.New(cron.WithSeconds()), //New returns a new Cron job runner, in the Local time zone.
		rmqChan: rmqChan,
		// outChan: ch,
	}
}

func (sched *Scheduler) Start() {
	// runs this function every minute
	sched.cron.AddFunc("0 * * * * *", sched.scheduleJob)
	sched.cron.Start()
}

func (sched *Scheduler) Stop() {
	fmt.Println("Stopping scheduler")
	sched.cron.Stop()
}

func (sched *Scheduler) scheduleJob() {
	fmt.Printf("Scanning for new order(s) at %v\n", time.Now().Format("2006-Jan-02 15:04:05"))
	resp, err := sched.getEmailForSending()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Scheduling %v email(s) at %v\n", len(resp), time.Now().Format("2006-Jan-02 15:04:05"))
	for _, em := range resp {
		body := em.String()
		err = sched.rmqChan.Publish(body)
		if err != nil {
			fmt.Println("Can't channel mes")
		}
	}
}

// getEmailForSending get email and fill up enough information ready for sending
func (sched *Scheduler) getEmailForSending() ([]*sendgrid.EmailContent, error) {
	resp, err := sched.prepareContent()
	if err != nil {
		return resp, err
	}
	// fill FromUser
	// why we can set FromUser here? aws:
	// I think Sender usualy only  uses a fixed mail or if we want divide the work for other admin mail
	for _, emailContent := range resp {
		emailContent.FromUser = &sendgrid.EmailUser{
			Name:  DefaultFromName,
			Email: DefaultFromEmail,
		}
	}

	return resp, err
}

// scanFromDB get all orders that match the predefined condition (created_at < now - 1 min && thankyou_email_sent == falses)
func (sched *Scheduler) prepareContent() ([]*sendgrid.EmailContent, error) {
	var resp []*sendgrid.EmailContent
	// Where i prepare before query
	startTime := time.Now().Add(-time.Hour * 20).Format("2006-01-02 15:04:05")
	fmt.Println("Timeeee ", startTime)
	endTime := time.Now().Format("2006-01-02 15:04:05")
	// Convert to string
	st := fmt.Sprintf("%v", startTime)
	et := fmt.Sprintf("%v", endTime)

	productReport, err := model.BestSellProducts(st, et)
	Error(err)
	orderReport, err := model.OrderAnalysis(st, et)
	Error(err)

	pathOrp, err := barchart.CreateChartOrder(orderReport)
	Error(err)
	pathPdr, err := barchart.CreateChartProduct(productReport)
	Error(err)
	paths := make([]string, 0)
	paths = append(paths, pathOrp, pathPdr)

	ats := make([]maller.Attachment, 0)
	for _, v := range paths {
		dat, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println(err)
		}
		encoded := base64.StdEncoding.EncodeToString([]byte(dat))
		at := maller.Attachment{
			Content:     encoded,
			Type:        "image/png",
			Filename:    "chart.png",
			Disposition: "inline",
			ContentID:   "ProductReport",
		}
		ats = append(ats, at)
	}
	fmt.Println("len", len(ats))

	resp = append(resp, &sendgrid.EmailContent{
		Subject:          DefaultReportSubject,
		PlainTextContent: DefaultReportBodyPlain,
		HtmlContent:      DefaultReportBodyHtml,
		Attachments:      ats,
		ToUser: &sendgrid.EmailUser{
			Name:  DefaultShopOwnerNamme,
			Email: DefaultShopOwnerEmail,
		},
	})
	return resp, nil
}

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
