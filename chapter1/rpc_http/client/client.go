package client

import (
	"fmt"
	"gopractice/microservicewithgo/chapter1/rpc_http/contract"
	"log"
	"net/rpc"
)

const port = 1234

func CreateClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("Client Create Error")
	}
	return client
}

func PerformRequest(c *rpc.Client) contract.HelloResponse {
	res := contract.HelloResponse{}
	req := contract.HelloRequest{Name: "tj"}
	c.Call("HelloWorldHandler.Hello", &req, &res)
	return res
}
