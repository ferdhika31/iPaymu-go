package ipaymu

type Customer struct {
	Name  string
	Phone string
	Email string
}

func NewCustomer(name string, phone string, email string) Customer {
	return Customer{
		Name:  name,
		Phone: phone,
		Email: email,
	}
}
