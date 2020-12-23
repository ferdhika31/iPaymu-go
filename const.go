package ipaymu

type PaymentMethodType int8

const (
	_ PaymentMethodType = iota

	VA
	BankTransfer
	ConvenienceStore
	COD
	PayLater
	QRIS
)

var typeStr = map[PaymentMethodType]string{
	VA:               "va",
	BankTransfer:     "banktransfer",
	ConvenienceStore: "cstore",
	COD:              "cod",
	PayLater:         "paylater",
	QRIS:             "qris",
}

func (pct PaymentMethodType) String() string {
	for idx, ts := range typeStr {
		if idx == pct {
			return ts
		}
	}
	return "null"
}
