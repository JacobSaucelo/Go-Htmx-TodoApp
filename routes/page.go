package todoapp

import (
	"fmt"
	"net/http"
)

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func RoutesMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route page")
}

func TodoFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
