package main

import (
	"fmt"
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
