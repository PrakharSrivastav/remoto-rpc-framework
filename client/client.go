package main

import (
	"context"
	"fmt"
	"github.com/PrakharSrivastav/remoto-rpc-framework/client/stub"
	"net/http"
	"time"
)

func main() {

	// create an http.Client
	clientHTTP := http.Client{
		Timeout:   2 * time.Second,
		Transport: &http.Transport{IdleConnTimeout: 5 * time.Second},
	}

	// create a new Bank client
	c := stub.NewBankClient("http://localhost:8080", &clientHTTP)
	ctx := context.Background()

	// call the Pay Method on the Bank client
	p, err := c.Pay(ctx, &stub.PaymentRequest{AccountID: "Some ID", Amount: 100})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", p)
}
