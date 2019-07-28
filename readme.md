# remoto-rpc-framework

This repository is written to support the blog post [here](https://www.prakharsrivastav.com/posts/remoto-rpc-framework/).



## Install Remoto

```bash
go get -v github.com/machinebox/remoto
```

## Download this repo and build

```bash
git clone https://github.com/PrakharSrivastav/remoto-rpc-framework.git
cd remoto-rpc-framework
go build ./...
```


## Generation scripts

```bash
# server code generation
remoto generate  schema/bank.remoto.go templates/client.go.plush -o client/stub/stub.go \
&& gofmt -w ./client/stub/stub.go

# client code generation
remoto generate  schema/bank.remoto.go templates/server.go.plush -o server/skeleton/service.go \
&& gofmt -w ./server/skeleton/service.go
```

## Run the server

from the project root, run:
```bash
prakhar@tardis (master)✗ % go run server/server.go  
starting server on 8080
endpoint: /remoto/Bank.Balance
endpoint: /remoto/Bank.Pay
endpoint: /remoto/Bank.WithDraw
```

## Run the client

from the project root, run:

```bash
prakhar@tardis (master)✗ [1] % go run client/client.go
&stub.Response{OK:true, Message:"Amount credited", Amount:100, Error:""} 
```