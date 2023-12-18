package main

import (
	"fmt"
	"net/http"
	todoapp "todoapp/routes"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", todoapp.RoutesMain)
}

func todo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
