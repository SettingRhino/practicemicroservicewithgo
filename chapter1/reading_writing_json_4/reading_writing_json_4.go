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
	cathandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))
	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Println("server Start")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// curl localhost:8080/helloworld -d '{"name":"tj"}'
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decorder := json.NewDecoder(r.Body)
	decorder.Decode(&request)
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }

	// var request helloWorldRequest
	// err = json.Unmarshal(body, &request)
	// if err != nil {
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }
	response := helloWorldResponse{Message: "hello" + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

}
