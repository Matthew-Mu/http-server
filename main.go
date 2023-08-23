package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	//connect to postgresql
	db, err := Setup()
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Connected")

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(Todo{})
	fmt.Println("Migrated")
	/*
		todo := Todo{
			Title:       "Complete Postgres integration",
			Description: "add full CRUD implementation",
			Status:      "Not_Complete",
		}

		result, err := CreateTodo(db, todo)
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println("Todo Created", result)
	*/
	todoArr := RetrieveAll(db)
	id := todoArr[1].ID.String()
	//id := idu.String()

	//fmt.Println("Input status: ")
	//fmt.Scanln(&id)
	/*
		retrievedTodo, _ := SelectPaymentWIthId(db, id)
		fmt.Println("Your TODO is \n", retrievedTodo.Title)
	*/
	updatedTodo, _ := UpdateTodo(db, id, Todo{
		Status: "Complete",
	})

	fmt.Println("STATUS = ", updatedTodo)

	payment, _ := SelectPaymentWIthId(db, id)
	fmt.Println("Your payment is", payment)

	// delete a payment with previous id
	//DeleteTodo(db, id)
	//fmt.Println("Your payment now is deleted")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Todo{
			"Todos": todoArr,
		}
		tmpl.Execute(w, films)
	}

	//fileserver := http.FileServer(http.Dir("./static"))

	//http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/template", h1)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
