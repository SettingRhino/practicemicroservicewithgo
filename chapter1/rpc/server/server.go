package server

import (
	"fmt"
	"gopractice/microservicewithgo/chapter1/rpc/contract"
	"log"
	"net"
	"net/rpc"
	"sync"
)

const port = 1234

func StartServer(wg *sync.WaitGroup) {
	helloWorld := &HelloWorld{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("port:%v, impossible RPC", port))
	}
	defer l.Close()
	wg.Done()
	for {
		con, _ := l.Accept()
		go rpc.ServeConn(con)
	}

}

type HelloWorld struct{}

func (h *HelloWorld) HelloWorldHandler(args *contract.HelloRequest, reply *contract.HelloResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
