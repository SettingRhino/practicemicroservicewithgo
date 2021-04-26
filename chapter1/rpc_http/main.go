package main

import (
	"fmt"
	_ "fmt"
	"gopractice/microservicewithgo/chapter1/rpc_http/client"
	"gopractice/microservicewithgo/chapter1/rpc_http/server"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go server.StartServer(&wg)
	wg.Wait()
	c := client.CreateClient()
	res := client.PerformRequest(c)
	fmt.Println(res)
}
