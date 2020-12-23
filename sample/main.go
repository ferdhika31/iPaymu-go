package main

import (
	"encoding/json"
	"fmt"

	ipaymu "github.com/ferdhika31/iPaymu-go"
	cnstipay "github.com/ferdhika31/iPaymu-go/consts"
)

var API_SECRET = "387AF695-EE0B-47BD-91A4-ADF55261159D"
var VA = "1179003821708289"

func redirPayment() {
	var cfg = ipaymu.NewConfig()
	cfg.Env = cnstipay.Sandbox
	cfg.ApiSecret = API_SECRET
	cfg.VirtualAccount = VA
	cfg.NotifyUrl = "http://localhost:8000/notify"
	cfg.ReturnUrl = "http://localhost:8000/return"
	cfg.CancelUrl = "http://localhost:8000/cancel"

	var payment = ipaymu.Payment{Config: cfg}
	// optional
	var c = ipaymu.NewCustomer("Ferdhika", "08313213131", "fer@dika.web.id")
	payment.SetCustomer(c)
	// mandatory
	payment.AddProduct(ipaymu.Product{Name: "Pot Kayu", Qty: 1, Price: 18000})
	payment.AddProduct(ipaymu.Product{Name: "Kaos", Qty: 2, Price: 60000})

	res, _ := payment.RedirectPayment(&ipaymu.RedirectRequest{
		// optional
		ReferenceID:   "TRX0031",
		PaymentMethod: cnstipay.VA,
	})

	// res, _ := payment.RedirectPayment(&ipaymu.RedirectRequest{})

	jsonString, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}

func getBalance() {
	var cfg = ipaymu.NewConfig()
	cfg.Env = cnstipay.Sandbox
	cfg.ApiSecret = API_SECRET

	b := ipaymu.Balance{Config: cfg}

	res, _ := b.GetBalance()

	jsonString, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}

func getTransaction(id uint32) {
	var cfg = ipaymu.NewConfig()
	cfg.Env = cnstipay.Sandbox
	cfg.ApiSecret = API_SECRET

	t := ipaymu.Transaction{Config: cfg}

	res, _ := t.GetTransaction(id)

	jsonString, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonString))
}

func main() {
	redirPayment()
	getBalance()
	getTransaction(30029)
}
