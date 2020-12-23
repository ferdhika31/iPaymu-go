package consts

type PaymentMethodType string

const (
	VA               PaymentMethodType = "va"
	BankTransfer     PaymentMethodType = "banktransfer"
	ConvenienceStore PaymentMethodType = "cstore"
	COD              PaymentMethodType = "cod"
	PayLater         PaymentMethodType = "paylater"
	QRIS             PaymentMethodType = "qris"
)
