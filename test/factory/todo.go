package factory

import (
	"github.com/api/internal/infra/models"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func (Factory) CreateTodo(db *gorm.DB) models.Todo {
	var todo models.Todo
	faker.FakeData(&todo)
	db.Create(&todo)
	return todo
}
