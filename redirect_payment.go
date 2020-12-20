package iPaymu

type RedirectPayment struct {
	Config   Config
	Products []Product
	Customer Customer
}

func (rp *RedirectPayment) SetCustomer(c Customer) Customer {
	rp.Customer = c
	return rp.Customer
}

func (rp *RedirectPayment) AddProduct(p Product) []Product {
	rp.Products = append(rp.Products, p)
	return rp.Products
}

func (rp *RedirectPayment) Create(body map[string]interface{}) map[string]interface{} {

	// set buyer info
	body["buyerName"] = rp.Customer.Name
	body["buyerEmail"] = rp.Customer.Email
	body["buyerPhone"] = rp.Customer.Phone

	// set notifyUrl
	if rp.Config.NotifyUrl != "" {
		body["notifyUrl"] = rp.Config.NotifyUrl
	}
	// set cancelUrl
	if rp.Config.CancelUrl != "" {
		body["cancelUrl"] = rp.Config.CancelUrl
	}
	// set returnUrl
	if rp.Config.ReturnUrl != "" {
		body["returnUrl"] = rp.Config.ReturnUrl
	}

	var products = []string{}
	var qtys = []uint32{}
	var prices = []uint32{}

	for _, p := range rp.Products {
		products = append(products, p.Name)
		qtys = append(qtys, p.Qty)
		prices = append(prices, p.Price)
	}

	body["product"] = products
	body["qty"] = qtys
	body["price"] = prices

	// fmt.Print(body)

	res := Request("POST", "/api/v2/payment", body, rp.Config)

	return res
}
