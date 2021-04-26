package client

import (
	"fmt"
	"gopractice/microservicewithgo/chapter1/rpc/contract"
	"log"
	"net/rpc"
)

const port = 1234

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func PerformRequest(rpcclient *rpc.Client) contract.HelloResponse {
	var reply contract.HelloResponse
	req := &contract.HelloRequest{Name: "tj"}
	err := rpcclient.Call("HelloWorld.HelloWorldHandler", req, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	return reply
}
