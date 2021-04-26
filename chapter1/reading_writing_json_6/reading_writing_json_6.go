package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080
	fileserver := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", fileserver))
	http.HandleFunc("/hellowWorlde", helloWorldHandler)
	http.Handle("/nourl", http.NotFoundHandler())
	log.Println("serverStart")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)

	response := helloWorldResponse{Message: "hello" + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
