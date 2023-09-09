package main

import (
	"fmt"
	"log"
	"net/http"
)


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
			Title:       "test HTMX integration",
			Description: "implement various HTMX handlers",
			Status:      "Not_Complete",
		}

		for i := 0; i < 10; i++ {
			result, err := CreateTodo(db, todo)
			if err != nil {
				log.Panic(err)
				return
			}
			fmt.Println("Todo Created", result)
		}*/
	//id := todoArr[1].ID.String()
	//id := idu.String()

	//fmt.Println("Input status: ")
	//fmt.Scanln(&id)
	/*
			retrievedTodo, _ := SelectPaymentWIthId(db, id)
			fmt.Println("Your TODO is \n", retrievedTodo.Title)

		updatedTodo, _ := UpdateTodo(db, id, Todo{
			Status: "Complete",
		})

		fmt.Println("STATUS = ", updatedTodo)
	*/
	//payment, _ := SelectPaymentWIthId(db, id)
	//fmt.Println("Your payment is", payment)

	// delete a payment with previous id
	//DeleteTodo(db, id)
	//fmt.Println("Your payment now is deleted")

	// handler function #1 - returns the index.html template, with film data

	//fileserver := http.FileServer(http.Dir("./static"))

	//http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", h1(db))
	http.HandleFunc("/add-film/", h2(db))
	http.HandleFunc("/delete-todo", deleteHandler(db))
	http.HandleFunc("/update-todo", updateHandler(db))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
