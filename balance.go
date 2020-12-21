package ipaymu

type Balance struct {
	Config Config
}

func (b *Balance) GetBalance() map[string]interface{} {
	res := Request("GET", "/api/saldo?key="+b.Config.ApiSecret, nil, b.Config)

	return res
}
