package main

import (
	//"context"
	"context"
	"fmt"

	"github.com/PrakharSrivastav/remoto-rpc-framework/server/skeleton"
)

func main() {
	addr := "localhost:8080"
	fmt.Println("starting server on 8080")
	if err := skeleton.Run(addr, service{}); err != nil {
		panic(err)
	}
}


// service will implement the Bank Service and provide the service side business logic
type service struct{}

// Balance checks the current balance by Account ID
func (service) Balance(context.Context, *skeleton.BalaceRequest) (*skeleton.Response, error) {
	fmt.Println("balance endpoint")
	return &skeleton.Response{OK: true, Message: "Request success", Amount: 100, Error: "",}, nil
}

// Pay credits to the account
func (service) Pay(context.Context, *skeleton.PaymentRequest) (*skeleton.Response, error) {
	fmt.Println("pay endpoint")
	return &skeleton.Response{OK: true, Message: "Amount credited", Amount: 100, Error: "",}, nil
}

// Withdraw debits the account
func (service) WithDraw(context.Context, *skeleton.WithdrawlRequest) (*skeleton.Response, error) {
	fmt.Println("withdraw endpoint")
	return &skeleton.Response{OK: true, Message: "Amount withdrawn", Amount: 100, Error: "",}, nil
}
