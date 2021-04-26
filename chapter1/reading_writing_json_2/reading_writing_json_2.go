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
type helloWorldResponseDynamicTag struct {
	// 필드의 태그값이용하기
	// 출력을 message로 한다
	Message string `json:"message"`
	// 출력하지 않는다
	Author string `json:"-"`
	// 없으면 출력하지 않는다
	Date string `json:",omitempty"`
	// 이름은 id로 그리고 string으로 출력한다.
	Id int `json:"id,string"`
}

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Println("serverStart")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// (ResponseWriter, *Request)
func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	// response := helloWorldResponse{Message: "hello World"}
	response := helloWorldResponseDynamicTag{Message: "hello World", Author: "Nooutput-data", Date: "2021-04-24", Id: 256}
	data, err := json.Marshal(response)
	// data, err := json.MarshalIndent(response, "prefix", "		")
	if err != nil {
		panic("Ooops!")
	}
	fmt.Fprint(w, string(data))
}
