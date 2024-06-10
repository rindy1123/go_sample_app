package models

import (
	"github.com/api/internal/infra"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoStatus string

const (
	TODO       TodoStatus = "todo"
	INPROGRESS TodoStatus = "inprogress"
	DONE       TodoStatus = "done"
)

type Todo struct {
	Base   infra.Base `gorm:"embedded" faker:"-"`
	Title  string     `gorm:"title"    faker:"word"`
	Status TodoStatus `gorm:"status"   faker:"oneof:todo,inprogress,done"`
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

func (todo Todo) Create(db *gorm.DB) (uuid.UUID, error) {
	result := db.Create(&todo)
	return todo.Base.ID, result.Error
}

func (todo Todo) Update(db *gorm.DB, id string) Todo {
	result := db.Model(&todo).Clauses(clause.Returning{}).Where("id = ?", id).Updates(todo)

	if result.Error != nil {
		panic(result.Error)
	}
	return todo
}

func (todo Todo) HardDelete(db *gorm.DB, id string) {
	db.Unscoped().Delete(&Todo{}, "id = ?", id)
}
