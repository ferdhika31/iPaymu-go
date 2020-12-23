package consts

type PaymentChannelType string

const (
	// va
	BAG     PaymentChannelType = "bag"
	BNI     PaymentChannelType = "bni"
	CIMB    PaymentChannelType = "cimb"
	Mandiri PaymentChannelType = "mandiri"
	// bank transfer
	BCA PaymentChannelType = "bca"
	// cstore
	Alfamart  PaymentChannelType = "alfamart"
	Indomaret PaymentChannelType = "indomaret"
	// cod
	RPX PaymentChannelType = "rpx"
	// paylater
	Akulaku PaymentChannelType = "akulaku"
	// qris
	Linkaja PaymentChannelType = "linkaja"
)
