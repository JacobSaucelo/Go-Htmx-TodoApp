package todoapp

import (
	"fmt"
	"html/template"
	"net/http"
)

var RoutesTmpl *template.Template

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
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{
				Item: "todo 1",
				Done: false,
			},
			{
				Item: "todo 2",
				Done: true,
			},
			{
				Item: "todo 3",
				Done: false,
			},
		},
	}
}
