package contract

type HelloResponse struct {
	Message string `json:"message"`
}

type HelloRequest struct {
	Name string `json:"name"`
}
