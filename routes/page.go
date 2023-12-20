package todoapp

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
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
	RoutesTmpl = template.Must(template.ParseFiles("pages/index.gohtml"))
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

	RoutesTmpl.Execute(w, data)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)

	item := r.PostFormValue("todo-details")

	RoutesTmpl = template.Must(template.ParseFiles("pages/index.gohtml"))
	RoutesTmpl.ExecuteTemplate(w, "list-todos", Todo{
		Item: item,
		Done: false,
	})
}
