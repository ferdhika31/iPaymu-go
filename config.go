package ipaymu

import "github.com/ferdhika31/iPaymu-go/consts"

type Config struct {
	Env            consts.EnvType
	ApiSecret      string
	VirtualAccount string
	NotifyUrl      string
	CancelUrl      string
	ReturnUrl      string
}

func NewConfig() Config {
	return Config{
		Env: consts.Production,
	}
}
