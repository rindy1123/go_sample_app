package models_test

import (
	"testing"

	"github.com/api/internal/models"
	"github.com/api/test/factory"
)

func TestTodo_Create(t *testing.T) {
	db := models.SetupTestDB()
	id, err := models.Todo{Title: "test", Status: "todo"}.Create(db)
	if err != nil {
		t.Fatalf(`Error: %v`, err)
	}
	var todo models.Todo
	db.Find(&todo, "id = ?", id)
	if todo.Title != "test" || todo.Status != "todo" {
		t.Fatalf(`Expected todo to be {Title: "test", Status: "todo"}, got %v`, todo)
	}
}

func TestTodo_Get(t *testing.T) {
	db := models.SetupTestDB()
	todo := factory.Factory{}.CreateTodo(db)
	got := models.Todo{}.Get(db, todo.Base.ID.String())
	if got.Title != todo.Title || got.Status != todo.Status {
		t.Fatalf(`Expected todo to be {Title: "test", Status: "todo"}, got %v`, todo)
	}
}

func TestTodo_List(t *testing.T) {
	db := models.SetupTestDB()
	factory.Factory{}.CreateTodo(db)
	factory.Factory{}.CreateTodo(db)
	todos := models.Todo{}.List(db)
	if len(todos) != 2 {
		t.Fatalf(`Expected 2 todos, got %v`, len(todos))
	}
}

func TestTodo_Update(t *testing.T) {
	db := models.SetupTestDB()
	todo := factory.Factory{}.CreateTodo(db)
	got := models.Todo{Title: "updated"}.Update(db, todo.Base.ID.String())
	if got.Title != "updated" || len(got.Status) == 0 {
		t.Fatalf(`Expected todo to be {Title: "updated", Status: "done"}, got %v`, got)
	}
}

func TestTodo_HardDelete(t *testing.T) {
	db := models.SetupTestDB()
	todo := factory.Factory{}.CreateTodo(db)
	models.Todo{}.HardDelete(db, todo.Base.ID.String())
	var got models.Todo
	db.Find(&got, "id = ?", todo.Base.ID)
	if got.Title != "" || got.Status != "" {
		t.Fatalf(`Expected todo to be {Title: "", Status: ""}, got %v`, got)
	}
}
