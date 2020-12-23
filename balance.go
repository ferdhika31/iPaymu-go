package ipaymu

type Balance struct {
	Config Config
}

func (b *Balance) GetBalance() Response {

	res, _ := Call("GET", "/api/saldo?key="+b.Config.ApiSecret, nil, b.Config)

	return res
}
