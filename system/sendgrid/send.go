package sendgrid

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Mailer defines function for sending email
type Mailer interface {
	Send(*EmailContent) error
}

// EmailUser defines email address info
type EmailUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// encoding to json

func (eu *EmailUser) String() string {
	b, _ := json.Marshal(eu)
	return string(b)
}

// EmailContent defines email content info
type EmailContent struct {
	ID               int64      `json:"id"`
	StartTime        string     `json:"start_time"`
	EndTime          string     `json:"enf_time"`
	Subject          string     `json:"subject"`
	FromUser         *EmailUser `json:"from"`
	ToUser           *EmailUser `json:"to"`
	PlainTextContent string     `json:"plaintext_content"`
	HtmlContent      string     `json:"html_content"`
	Files            []string   `json:"attachments"`
}

// encoding mail to json
func (mailCt *EmailContent) String() string {
	b, _ := json.Marshal(mailCt)
	return string(b)
}

// Validate will check whether the email content is valid
func (mailCt *EmailContent) Validate() error {
	if mailCt == nil || mailCt.FromUser == nil || mailCt.ToUser == nil || mailCt.PlainTextContent == "" {
		return errors.New("wrong content")
	}
	return nil
}

// new client NewSendgrid creates new Sendgrid client
func NewSendgrid(apiKey string) *Sendgrid {
	client := sendgrid.NewSendClient(apiKey)
	return &Sendgrid{
		Client: client,
	}
}

// Sendgrid implements logic to send email to destination(location/ diem den) email address via sendgrid
type Sendgrid struct {
	Client *sendgrid.Client
}

// Send will send email based on email content
func (sm *Sendgrid) Send(mailCt *EmailContent) error {
	m := mail.NewV3Mail()
	if err := mailCt.Validate(); err != nil {
		return err
	}
	// Host infrom
	from := mail.NewEmail(mailCt.FromUser.Name, mailCt.FromUser.Email)
	to := mail.NewEmail(mailCt.ToUser.Name, mailCt.ToUser.Email)
	plainTextContent := mail.NewContent("text/plain", mailCt.PlainTextContent)
	m.SetFrom(from)
	m.AddContent(plainTextContent)
	// Reciever info
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	date := time.Now().Format("02-01-2006")
	personalization.Subject = "Report - Daily! " + date
	m.AddPersonalizations(personalization)

	// Product chart
	pdPng, err := ioutil.ReadFile(mailCt.Files[0])
	if err != nil {
		fmt.Println(err)
	}
	encodedPdc := base64.StdEncoding.EncodeToString([]byte(pdPng))
	prt := &mail.Attachment{
		Content:     encodedPdc,
		Type:        "image/png",
		Filename:    "Product.png",
		Disposition: "inline",
		ContentID:   "ProductReport",
	}

	// Order chart
	orderchart, err := ioutil.ReadFile(mailCt.Files[1])
	if err != nil {
		return err
	}
	encodeOdc := base64.StdEncoding.EncodeToString([]byte(orderchart))
	ort := &mail.Attachment{
		Content:     encodeOdc,
		Type:        "image/png",
		Filename:    "Order.png",
		Disposition: "inline",
		ContentID:   "ProductReport",
	}

	m.AddAttachment(prt)
	m.AddAttachment(ort)
	response, err := sm.Client.Send(m)
	if err != nil {
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		fmt.Println("Sending email: ", mailCt)
		//  Uncomment to remove the file after sent
		// e := os.Remove(mailCt.Files[1])
		// if e != nil {
		// 	log.Fatal(e)
		// }
	}
	return nil
}
