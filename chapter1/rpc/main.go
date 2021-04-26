package main

import (
	"fmt"
	"sync"

	"gopractice/microservicewithgo/chapter1/rpc/client"

	"gopractice/microservicewithgo/chapter1/rpc/server"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go server.StartServer(&wg)
	wg.Wait()
	c := client.CreateClient()
	defer c.Close()
	reply := client.PerformRequest(c)
	fmt.Println(reply)
}
