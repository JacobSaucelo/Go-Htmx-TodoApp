package main

import (
	"html/template"
	"net/http"
	todoapp "todoapp/routes"
)

func main() {
	mux := http.NewServeMux()
	todoapp.RoutesTmpl = template.Must(template.ParseFiles("pages/index.html"))

	mux.HandleFunc("/", todoapp.RoutesMain)
}
