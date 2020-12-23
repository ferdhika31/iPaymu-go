package ipaymu

import "github.com/ferdhika31/iPaymu-go/consts"

type RedirectRequest struct {
	ReferenceID   string                   `json:"referenceId,omitempty"`
	PaymentMethod consts.PaymentMethodType `json:"paymentMethod,omitempty"`
}

type DirectRequest struct {
	Amount         uint64                    `json:"amount,omitempty"`
	Expired        int64                     `json:"expired,omitempty"`
	ExpiredType    consts.ExpiredType        `json:"expiredType,omitempty"`
	Comments       string                    `json:"comments,omitempty"`
	ReferenceID    string                    `json:"referenceId,omitempty"`
	PaymentMethod  consts.PaymentMethodType  `json:"paymentMethod,omitempty"`
	PaymentChannel consts.PaymentChannelType `json:"paymentChannel,omitempty"`
}

type CashOnDelivery struct {
	Weight        int64
	Dimension     []int64
	PickupArea    string
	PickupAddress string
}
