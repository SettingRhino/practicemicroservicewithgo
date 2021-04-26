package server

import (
	"fmt"
	"gopractice/microservicewithgo/chapter1/rpc_http/contract"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

const port = 1234

func StartServer(wg *sync.WaitGroup) {
	var helloWorldHandler HelloWorldHandler
	rpc.Register(&helloWorldHandler)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("port:%v, impossible RPC", port))
	}
	defer l.Close()
	wg.Done()
	http.Serve(l, nil)

}

type HelloWorldHandler struct {
}

func (h *HelloWorldHandler) Hello(req *contract.HelloRequest, res *contract.HelloResponse) error {
	res.Message = "Hello" + req.Name
	return nil
}
