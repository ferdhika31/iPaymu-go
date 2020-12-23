package ipaymu

type RedirectRequest struct {
	ReferenceID   string            `json:"referenceId,omitempty"`
	PaymentMethod PaymentMethodType `json:"paymentMethod,omitempty"`
}

type CashOnDelivery struct {
	Weight        int64
	Dimension     []int64
	PickupArea    string
	PickupAddress string
}
