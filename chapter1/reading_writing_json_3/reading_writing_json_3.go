package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	Auth    string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id,string"`
}

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Println("server Start")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(*r)
	response := helloWorldResponse{Message: "hello", Auth: "", Date: "2020", Id: 586}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
