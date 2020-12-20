package main

import (
	"encoding/json"
	"fmt"

	"github.com/ferdhika31/iPaymu-go"
)

var API_SECRET = "387AF695-EE0B-47BD-91A4-ADF55261159D"
var VA = "1179003821708289"

func redirPayment() {
	var cfg = iPaymu.NewConfig()
	cfg.ApiSecret = API_SECRET
	cfg.VirtualAccount = VA
	cfg.NotifyUrl = "http://localhost:8000/notify"
	cfg.ReturnUrl = "http://localhost:8000/return"
	cfg.CancelUrl = "http://localhost:8000/cancel"
	rp := iPaymu.RedirectPayment{Config: cfg}
	// optional
	rp.SetCustomer(iPaymu.Customer{Name: "Ferdhika", Phone: "08392313213", Email: "fer@dika.web.id"})
	// mandatory
	rp.AddProduct(iPaymu.Product{Name: "Pot Kayu", Qty: 1, Price: 18000})
	rp.AddProduct(iPaymu.Product{Name: "Kaos", Qty: 2, Price: 60000})

	// optional
	payload := map[string]interface{}{
		"referenceId":   "TRX0031",
		"paymentMethod": "qris",
	}

	res := rp.Create(payload)

	jsonString, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}

func getBalance() {
	var cfg = iPaymu.NewConfig()
	cfg.ApiSecret = API_SECRET

	b := iPaymu.Balance{Config: cfg}

	res := b.GetBalance()

	jsonString, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}

func main() {
	redirPayment()
	getBalance()
}
