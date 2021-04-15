package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CustomResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/", rootHandler)
	_ = http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start handler request")

	queryForm, err := url.ParseQuery(r.URL.RawQuery)

	w.Header().Set("Content-Type", "application/json")
	message := ""

	if err == nil && len(queryForm["message"]) > 0 {
		message = queryForm["message"][0]
	} else {
		message = "Hello Go Server"
	}

	_ = json.NewEncoder(w).Encode(CustomResponse{200, message})
	fmt.Println("Handler request completed")
}
