package ipaymu

type Config struct {
	Env            string
	ApiSecret      string
	VirtualAccount string
	NotifyUrl      string
	CancelUrl      string
	ReturnUrl      string
}

func NewConfig() Config {
	return Config{
		Env: "Sandbox",
	}
}
