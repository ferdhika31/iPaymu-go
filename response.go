package ipaymu

type Response struct {
	Status int64
	// v1
	Keterangan string `json:"Keterangan,omitempty"`
	// v1 saldo
	Username        string `json:"Username,omitempty"`
	MerchantBalance int64  `json:"MerchantBalance,omitempty"`
	MemberBalance   int64  `json:"MemberBalance,omitempty"`
	Saldo           int64  `json:"Saldo,omitempty"`
	// v1 transaction detail
	ID          int64  `json:"ID,omitempty"`
	SessionID   string `json:"SessionId,omitempty"`
	ReferenceID string `json:"ReferenceId,omitempty"`
	Pengirim    string `json:"Pengirim,omitempty"`
	Penerima    string `json:"Penerima,omitempty"`
	Nominal     string `json:"Nominal,omitempty"`
	Biaya       string `json:"Biaya,omitempty"`
	Waktu       string `json:"Waktu,omitempty"`
	WaktuBayar  string `json:"WaktuBayar,omitempty"`
	Expired     string `json:"Expired,omitempty"`
	Tipe        string `json:"Tipe,omitempty"`
	VA          string `json:"Va,omitempty"`

	// v2 payment
	Message string        `json:"Message,omitempty"`
	Data    *ResponseData `json:"Data,omitempty"`
}

type ResponseData struct {
	SessionId     string `json:"SessionId,omitempty"`
	TransactionId int64  `json:"TransactionId,omitempty"`
	ReferenceId   string `json:"ReferenceId,omitempty"`
	Via           string `json:"Via,omitempty"`
	Channel       string `json:"Channel,omitempty"`
	PaymentNo     string `json:"PaymentNo,omitempty"`
	PaymentName   string `json:"PaymentName,omitempty"`
	Total         string `json:"Total,omitempty"`
	Fee           int64  `json:"Fee,omitempty"`
	Expired       string `json:"Expired,omitempty"`
	Note          string `json:"Note,omitempty"`
	Url           string `json:"Url,omitempty"`
}
