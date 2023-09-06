package main

import (
	"fmt"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Println("Deez Nuts")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request, todo_list []Todo) {
	tmpl := template.Must(template.ParseFiles("template.html"))
	todos := map[string][]Todo{
		"Todos": todo_list,
	}
	println("accessing template")
	tmpl.Execute(w, todos)

}

// handler function #2 - returns the template block with the newly added film, as an HTMX response
func h2(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		//htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		//tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		new_todo := Todo{Title: title, Description: director, Status: "Not_Complete"}
		add_todo := CreateTodo(db, new_todo)
		returned_todo, _ := SelectTodoByID(db, add_todo.String())
		tmpl.ExecuteTemplate(w, "film-list-element", returned_todo)
	}

}

func h1(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todoArr := RetrieveAll(db)
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		films := map[string][]Todo{
			"Todos": todoArr,
		}
		tmpl.Execute(w, films)
	}
}

func deleteHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PostFormValue("id")
		DeleteTodo(db, id)

		todoArr := RetrieveAll(db)
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		todos := map[string][]Todo{
			"Todos": todoArr,
		}
		fmt.Println(todos)
		tmpl.Execute(w, todos)
	}
}
