package main

import (
	rpc "github.com/YouDad/learn.go/net/rpc"
)

func main() {
	go rpc.StartServer()
	rpc.StartClient()
}
