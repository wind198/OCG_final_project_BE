package controller

import (
	"backend/api/model"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type OrderIDReceived struct {
	OrderID int
}
type PaymentIntentOutputStruct struct {
	Err                    error
	PaymentIntentSecretKey string
}

func writeOutputJson(outputStruct PaymentIntentOutputStruct, w http.ResponseWriter) {
	jsonOutput, err := json.Marshal(outputStruct)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonOutput)
}

func checkInternalError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Println(err)
		http.Error(w, "Unexpected Error occured. Please come back later.", http.StatusMethodNotAllowed)
		return true
	}
	return false
}

func HandlePayment(w http.ResponseWriter, r *http.Request) {
	log.Println("handling payment", r.Body)
	if r.Method != "POST" {
		http.Error(w, "Unexpected Error occured. Please come back later.", http.StatusMethodNotAllowed)
		return
	}

	var orderReceived OrderIDReceived
	err := json.NewDecoder(r.Body).Decode(&orderReceived)
	if checkInternalError(err, w) {
		return
	}

	orderFromDB, err := model.GetOneOrder(strconv.Itoa(orderReceived.OrderID))
	if checkInternalError(err, w) {
		return
	}
	totalBillToPay := orderFromDB.TotalPrice
	totalBillToPayInCent := int64(totalBillToPay * 100)
	log.Println("totalBillToPayInCent", totalBillToPay)
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(totalBillToPayInCent),
		Currency: stripe.String("usd"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}
	pi, err := paymentintent.New(params)
	if checkInternalError(err, w) {
		return
	}
	paymentintentOutput := PaymentIntentOutputStruct{
		Err:                    nil,
		PaymentIntentSecretKey: pi.ClientSecret,
	}
	log.Println(paymentintentOutput)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	writeOutputJson(paymentintentOutput, w)
}
