package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", todo)
}

func todo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
