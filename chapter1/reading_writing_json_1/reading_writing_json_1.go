package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Mesaage string
}

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	fmt.Println("Service start")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Mesaage: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}
	fmt.Fprintln(w, string(data))
}
