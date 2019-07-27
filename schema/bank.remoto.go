package bank

type Bank interface {
	Pay(PaymentRequest) Response
	WithDraw(WithdrawlRequest) Response
	Balance(BalaceRequest) Response
}

type PaymentRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type Response struct {
	OK      bool    `json:"ok"`
	Message string  `json:"message"`
	Amount  float64 `json:"amount"`
}

type WithdrawlRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type BalaceRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

//remoto generate example.remoto.go ../../templates/remotohttp/server.go.plush -o ./server/greeter/server.go && gofmt -w ./server/greeter/server.go
//remoto generate example.remoto.go ../../templates/remotohttp/client.go.plush -o ./client/greeter/client.go && gofmt -w ./client/greeter/client.go
