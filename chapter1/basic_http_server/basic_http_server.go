package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("server start!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
