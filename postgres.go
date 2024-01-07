package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	Title       string
	Description string
	Status      string
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time `gorm:"autoUpdateTime"`
}

/*
	func connectDB() {
		dsn := "host=localhost user=postgres password='' dbname=postgres port=5432 sslmode=disable TimeZone=Australia/Sydney"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

}
*/
const (
	Host     = "127.0.0.1"
	User     = "postgres"
	Password = "1"
	Name     = "go-todo"
	Port     = "5432"
)

func Setup() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Host,
		Port,
		User,
		Name,
		Password,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTodo(db *gorm.DB, todo Todo) uuid.UUID {

	result := db.Create(&todo)
	fmt.Println(result.Row())
	return todo.ID

}

func RetrieveAll(db *gorm.DB) []Todo {

	var todos []Todo
	db.Find(&todos)
	for i := 0; i < len(todos); i++ {
		fmt.Println("Todo: ", todos[i])
	}
	return todos

}

func SelectTodoByID(db *gorm.DB, id string) (Todo, error) {

	var todos []Todo
	db.Find(&todos)
	for i := 0; i < len(todos); i++ {
		fmt.Println("Todo Array:", todos[i])
	}

	var todo Todo
	result := db.First(&todo, "ID = ?", id)
	if result.RowsAffected == 0 {
		return Todo{}, errors.New("payment data not found")
	}
	return todo, nil
}

func UpdateTodo(db *gorm.DB, id string) (Todo, error) {
	var updateTodo Todo
	todo, _ := SelectTodoByID(db, id)
	todo.Status = "Complete"
	result := db.Model(&updateTodo).Where("id = ?", id).Updates(todo)
	if result.RowsAffected == 0 {
		return Todo{}, errors.New("payment data not update")
	}
	return updateTodo, nil
}
func DeleteTodo(db *gorm.DB, id string) (int64, error) {
	var deletedTodo Todo
	result := db.Where("id = ?", id).Delete(&deletedTodo)
	if result.RowsAffected == 0 {
		return 0, errors.New("Todo not deleted!")
	}
	return result.RowsAffected, nil
}
