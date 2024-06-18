package todos

import (
	"github.com/api/internal/infra/models"
	"github.com/google/uuid"
)

type GetTodo struct {
	ID     uuid.UUID         `json:"id"`
	Title  string            `json:"title"`
	Status models.TodoStatus `json:"status"`
}

type ListTodos struct {
	Todos []GetTodo `json:"todos"`
}

func toSingular(todo models.Todo) GetTodo {
	return GetTodo{
		ID:     todo.Base.ID,
		Title:  todo.Title,
		Status: todo.Status,
	}
}

func toMultiple(todos []models.Todo) ListTodos {
	getTodos := []GetTodo{}
	for _, todo := range todos {
		getTodos = append(getTodos, toSingular(todo))
	}
	return ListTodos{Todos: getTodos}
}
