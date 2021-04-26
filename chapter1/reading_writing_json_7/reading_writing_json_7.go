package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type validContextString string

func main() {
	port := 8080
	validhandler := newValidHandler(newRealHandle())
	http.Handle("/helloworld", validhandler)
	log.Println("server Start")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloResponse struct {
	Message string `json:"message"`
}

type helloRequest struct {
	Name string `json:"name"`
}

type validHandler struct {
	next http.Handler
}

func newValidHandler(next http.Handler) http.Handler {
	return validHandler{next: next}
}

func (vh validHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := helloRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	c := context.WithValue(r.Context(), validContextString("name"), request.Name)
	r = r.WithContext(c)

	vh.next.ServeHTTP(w, r)
}

func newRealHandle() http.Handler {
	return realHnadler{}
}

type realHnadler struct {
}

func (rh realHnadler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validContextString("name")).(string)
	response := helloResponse{Message: "hello" + name}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
