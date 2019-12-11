package rpc

import (
	"fmt"
	"log"
	"net/rpc"
	"reflect"
)

func StartClient() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	var reply int
	args := &Args{7, 8}

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	args = &Args{77, 8}
	quotient := new(Quotient)

	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	result := replyCall.Reply.(*Quotient)

	// check errors, print, etc.
	fmt.Printf("Type: %s\n", reflect.TypeOf(replyCall.Reply).String())
	fmt.Printf("Arith: %d/%d=%d rem %d\n", args.A, args.B, result.Quo, result.Rem)

	client.Close()
}
