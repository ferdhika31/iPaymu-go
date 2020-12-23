package ipaymu

import "fmt"

type Transaction struct {
	Config Config
}

func (t *Transaction) GetTransaction(ids uint32) Response {

	id := fmt.Sprintf("%d", ids)

	res, _ := Call("GET", "/api/transaksi?key="+t.Config.ApiSecret+"&id="+id+"&format=json", nil, t.Config)

	return res
}
