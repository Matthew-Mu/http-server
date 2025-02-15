package main

import (
	"Matthew-Mu/http-server/weather"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type Environmentals struct {
	Pressure    string
	Temperature string
}

type Data struct {
	Todos   []Todo
	Weather []weather.Weather
}

func getWeatherTableHandler(w http.ResponseWriter, r *http.Request) {
	bytes := weather.Fetch()
	wTable := weather.ConvertBytesToJson(bytes)
	tmpl := template.Must(template.ParseFiles("static/table.html"))
	wthr := map[string][]weather.Weather{
		"Weather": wTable.Table,
	}
	tmpl.ExecuteTemplate(w, "WeatherTable", wthr)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	//fmt.Println("Deez Nuts")
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
	//println("accessing template")
	tmpl.Execute(w, todos)
}

// handler function addHandler - returns the template block with the newly added todo, as an HTMX response
func addHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("static/film-list-tmpl.html"))
		new_todo := Todo{Title: title, Description: director, Status: "Not_Complete"}
		CreateTodo(db, new_todo)
		todoArr := RetrieveAll(db)
		todos := map[string][]Todo{
			"Todos": todoArr,
		}
		//tmpl.ExecuteTemplate(w, "film-list-element", returned_todo)
		tmpl.Execute(w, todos)
	}
}

func homeHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todoArr := RetrieveAll(db)
		tmpl := template.Must(template.ParseFiles("static/index.html", "static/table.html", "static/header.html"))
		var todos = make(map[string][]Todo)
		todos["Todos"] = todoArr
		tmpl.ExecuteTemplate(w, "index", todos)
	}
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/table.html", "static/weather.html", "static/header.html"))
	bytes := weather.Fetch()
	wTable := weather.ConvertBytesToJson(bytes)
	wthr := map[string][]weather.Weather{
		"Weather": wTable.Table,
	}
	tmpl.ExecuteTemplate(w, "weather", wthr)
}

func deleteHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PostFormValue("id")
		DeleteTodo(db, id)
		todoArr := RetrieveAll(db)
		tmpl := template.Must(template.ParseFiles("static/film-list-tmpl.html"))
		todos := map[string][]Todo{
			"Todos": todoArr,
		}
		tmpl.Execute(w, todos)
	}
}

func updateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PostFormValue("id")
		UpdateTodo(db, id)
		todoArr := RetrieveAll(db)
		tmpl := template.Must(template.ParseFiles("static/update-tmpl.html"))
		todos := map[string][]Todo{
			"Todos": todoArr,
		}
		//fmt.Println(todos)
		tmpl.Execute(w, todos)
	}
}
func enviroHandler(w http.ResponseWriter, r *http.Request) {

	var enviro Environmentals
	res, err := http.Get("http://Pi5-1/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = json.Unmarshal(resData, &enviro)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(enviro.Pressure)
	tmpl := template.Must(template.ParseFiles("static/env.html", "static/index.html", "static/header.html"))
	/*_ := map[string]Environmentals{
		"env": enviro,
	}*/
	tmpl.ExecuteTemplate(w, "Environmentals", enviro)
}
