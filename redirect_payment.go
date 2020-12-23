package ipaymu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// create redirect payment
func (p *Payment) RedirectPayment(rr *RedirectRequest) (Response, error) {
	req := make(map[string]interface{})

	// req = structs.Map(rr)
	if rr.PaymentMethod != 0 {
		req["paymentMethod"] = fmt.Sprintf("%s", rr.PaymentMethod)
	}

	if rr.ReferenceID != "" {
		req["referenceId"] = fmt.Sprintf("%s", rr.ReferenceID)
	}

	// set buyer info
	req["buyerName"] = p.Customer.Name
	req["buyerEmail"] = p.Customer.Email
	req["buyerPhone"] = p.Customer.Phone

	// set notifyUrl
	if p.Config.NotifyUrl != "" {
		req["notifyUrl"] = p.Config.NotifyUrl
	}
	// set cancelUrl
	if p.Config.CancelUrl != "" {
		req["cancelUrl"] = p.Config.CancelUrl
	}
	// set returnUrl
	if p.Config.ReturnUrl != "" {
		req["returnUrl"] = p.Config.ReturnUrl
	}

	// set product list
	var products = []string{}
	var qtys = []uint32{}
	var prices = []uint32{}
	for _, prod := range p.Products {
		products = append(products, prod.Name)
		qtys = append(qtys, prod.Qty)
		prices = append(prices, prod.Price)
	}
	req["product"] = products
	req["qty"] = qtys
	req["price"] = prices

	var jsonBody []byte

	jsonBody, err := json.Marshal(req)

	if err != nil {
		log.Println(err)
	}

	// log.Println("Request Body : " + req)

	result, err := Call("POST", "/api/v2/payment", bytes.NewBuffer(jsonBody), p.Config)

	return result, err
}
