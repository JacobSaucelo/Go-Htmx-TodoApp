package main

import (
	"log"
	"net/http"
	todoapp "todoapp/routes"
)

func main() {
	mux := http.NewServeMux()
	// todoapp.RoutesTmpl = template.Must(template.ParseFiles("pages/index.gohtml"))

	mux.HandleFunc("/", todoapp.RoutesMain)
	mux.HandleFunc("/todo", todoapp.TodoFunc)

	log.Fatal(http.ListenAndServe(":6420", mux))
}
