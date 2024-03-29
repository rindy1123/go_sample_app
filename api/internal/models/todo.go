package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoStatus string

const (
	TODO       TodoStatus = "todo"
	INPROGRESS TodoStatus = "inprogress"
	DONE       TodoStatus = "done"
)

type Todo struct {
	Base   base       `gorm:"embedded"`
	Title  string     `gorm:"title"`
	Status TodoStatus `gorm:"status"`
}

func (Todo) List(db *gorm.DB) []Todo {
	var todos []Todo
	result := db.Find(&todos)

	if result.Error != nil {
		panic(result.Error)
	}
	return todos
}

func (Todo) Get(db *gorm.DB, id string) Todo {
	var todo Todo
	result := db.First(&todo, "id = ?", id)

	if result.Error != nil {
		panic(result.Error)
	}
	return todo
}

func (todo Todo) Create(db *gorm.DB) uuid.UUID {
	result := db.Create(&todo)

	if result.Error != nil {
		panic(result.Error)
	}
	return todo.Base.ID
}

func (todo Todo) Update(db *gorm.DB, id string) {
	result := db.Model(&todo).Where("id = ?", id).Updates(todo)

	if result.Error != nil {
		panic(result.Error)
	}
}

func (todo Todo) HardDelete(db *gorm.DB, id string) {
	db.Unscoped().Delete(&Todo{}, "id = ?", id)
}
