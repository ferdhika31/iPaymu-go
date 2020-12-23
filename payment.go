package ipaymu

type Payment struct {
	Config   Config
	Products []Product
	Customer Customer
}

func (p *Payment) SetCustomer(c Customer) Customer {
	p.Customer = c
	return p.Customer
}

func (p *Payment) AddProduct(prd Product) []Product {
	p.Products = append(p.Products, prd)
	return p.Products
}
