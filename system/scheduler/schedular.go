package scheduler

import (
	"OCG_final_project_BE/api/model"
	"OCG_final_project_BE/system/barchart"
	"OCG_final_project_BE/system/rbmq"
	"OCG_final_project_BE/system/sendgrid"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

/*
Goal:
 - Scan order, porduct and create chart for each
 - run at every 12 AM
 - after sent email update orders.report_sent

*/

const (
	DefaultReportSubject  = "Report for shop owner from Decorshop"
	DefaultReportBodyHtml = "<strong>Report for shop owner from Decorshop. Here's your reports today:</strong>"
	DefaultFromName       = "My Decor Shop"
	DefaultFromEmail      = "duynvwork@gmail.com"
	DefaultShopOwnerEmail = "ngoduy9911@gmail.com"
	DefaultShopOwnerNamme = "Duynv Decor"
)

type Scheduler struct {
	cron    *cron.Cron
	rmqChan *rbmq.Rabbit
	ctx     context.Context
}

func NewScheduler(ctx context.Context, rmqChan *rbmq.Rabbit) *Scheduler {
	return &Scheduler{
		ctx:     ctx,
		cron:    cron.New(cron.WithSeconds()), //New returns a new Cron job runner, in the Local time zone.
		rmqChan: rmqChan,
	}
}

func (sched *Scheduler) Start() {
	// second/minute/hour/day/month/dayweek
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
	// if we want to improve it we can
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
	startTime := time.Now().Add(-time.Hour * 300).Format("2006-01-02 15:04:05")
	endTime := time.Now().Format("2006-01-02 15:04:05")

	productReport, err := model.BestSellProducts(startTime, endTime)
	if err != nil {
		return resp, errors.New("can can get  best sellings product")
	}
	orderReport, err := model.OrderAnalysis(startTime, endTime)
	if err != nil {
		return resp, errors.New("c an not can not get order analysis")
	}

	// create chart and return file paths
	ordChart, err := barchart.CreateChartOrder(orderReport)
	if err != nil {
		return resp, errors.New("error create chart order")
	}
	pdtChart, err := barchart.CreateChartProduct(productReport)
	if err != nil {
		return resp, errors.New("error create chart product")
	}
	fs := []string{}
	fs = append(fs, pdtChart, ordChart)
	resp = append(resp, &sendgrid.EmailContent{
		StartTime:        startTime,
		EndTime:          endTime,
		Subject:          DefaultReportSubject,
		PlainTextContent: fmt.Sprintf("Today we got %v order(s) and total sales is %v, you are smart to come up with a marketing plan! ", orderReport.TotalOrders, orderReport.TotalSales),
		HtmlContent:      DefaultReportBodyHtml,
		Files:            fs,
		ToUser: &sendgrid.EmailUser{
			Name:  DefaultShopOwnerNamme,
			Email: DefaultShopOwnerEmail,
		},
	})
	return resp, nil
}
