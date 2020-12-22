package ipaymu

type Product struct {
	Name        string
	Qty         uint32
	Price       uint32
	Description string
}

type Products []Product

func NewProduct(name string, qty uint32, price uint32, desc string) Product {
	// add validation cuy

	return Product{
		Name:        name,
		Qty:         qty,
		Price:       price,
		Description: desc,
	}
}
