iPaymu REST API Client Golang
==============

Unofficial [iPaymu](https://ipaymu.com) API wrapper written in Golang for access from applications.

## Documentation

For the API documentation, please check [iPaymu API Documentation](https://ipaymu.com/en/api-documentation/).

## Feature

- [x] Redirect Payment
- [ ] Direct Payment
- [x] Get Balance Account
- [x] Detail Transaction

## To-do
- [ ] Clean Code
- [ ] Testing
- [ ] COD Feature in Direct and Redirect Payment

## Installation

```bash
go get -u github.com/ferdhika31/iPaymu-go
```

## Configuration

### Set Environment
By default the environment set the `production`. You can set manually if you need.

```go
import (
    ipaymu "github.com/ferdhika31/iPaymu-go"
    cnstipay "github.com/ferdhika31/iPaymu-go/consts"
)
var cfg = ipaymu.NewConfig()
cfg.Env = cnstipay.Sandbox
// or
cfg.Env = cnstipay.Production
```

### Set API Secret and Virtual Account

```go
cfg.ApiSecret = "API_SECRET_KEY_HERE"
cfg.VirtualAccount = "VA_KEY_HERE"
```

### Set URL

By default the `notifyUrl` is mandatory for `redirect payment` or `direct payment`. You can set url manually if you need. `returnUrl` and `cancelUrl` can only be used when `redirect payment`.

```go
cfg.NotifyUrl = "http://localhost:8000/notify"
cfg.ReturnUrl = "http://localhost:8000/return"
cfg.CancelUrl = "http://localhost:8000/cancel"
```

## Redirect Payment

```go
var payment = ipaymu.Payment{Config: cfg}
// optional
var c = ipaymu.NewCustomer("Ferdhika", "08313213131", "fer@dika.web.id")
payment.SetCustomer(c)
// mandatory
payment.AddProduct(ipaymu.Product{Name: "Pot Kayu", Qty: 1, Price: 18000})
payment.AddProduct(ipaymu.Product{Name: "Kaos", Qty: 2, Price: 60000})

res, _ := payment.RedirectPayment(&ipaymu.RedirectRequest{
    // optional
    ReferenceID:   "TRX0031",
    PaymentMethod: cnstipay.VA,
})
```

## Get Balance

```go
b := ipaymu.Balance{Config: cfg}
res, _ := b.GetBalance()
```

## Get Transaction Detail

```go
trx := ipaymu.Transaction{Config: cfg}
res, _ := trx.GetTransaction(id)
```

## Contributing

For any requests, bugs, or comments, please open an [issue](https://github.com/ferdhika31/iPaymu-go/issues) or [submit a pull request](https://github.com/ferdhika31/iPaymu-go/pulls).