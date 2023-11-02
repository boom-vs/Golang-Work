package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Strat Project Golang")
}
func main() {
	http.HandleFunc("/", handler)

	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
