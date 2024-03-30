package models_test

import (
	"testing"

	"github.com/api/internal/models"
	"github.com/api/test/factory"
)

func TestTodo_Create(t *testing.T) {
	db, teardown := models.SetupTestDB()
	defer teardown()
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
	db, teardown := models.SetupTestDB()
	defer teardown()
	todo := factory.Factory{}.CreateTodo(db)
	get := models.Todo{}.Get(db, todo.Base.ID.String())
	if get.Title != todo.Title || get.Status != todo.Status {
		t.Fatalf(`Expected todo to be {Title: "test", Status: "todo"}, got %v`, todo)
	}
}

func TestTodo_List(t *testing.T) {
	db, teardown := models.SetupTestDB()
	defer teardown()
	factory.Factory{}.CreateTodo(db)
	factory.Factory{}.CreateTodo(db)
	todos := models.Todo{}.List(db)
	if len(todos) != 2 {
		t.Fatalf(`Expected 2 todos, got %v`, len(todos))
	}
}

func TestTodo_Update(t *testing.T) {
	db, teardown := models.SetupTestDB()
	defer teardown()
	todo := factory.Factory{}.CreateTodo(db)
	models.Todo{Title: "updated", Status: "done"}.Update(db, todo.Base.ID.String())
	var get models.Todo
	db.Find(&get, "id = ?", todo.Base.ID)
	if get.Title != "updated" || get.Status != "done" {
		t.Fatalf(`Expected todo to be {Title: "updated", Status: "done"}, got %v`, get)
	}
}

func TestTodo_HardDelete(t *testing.T) {
	db, teardown := models.SetupTestDB()
	defer teardown()
	todo := factory.Factory{}.CreateTodo(db)
	models.Todo{}.HardDelete(db, todo.Base.ID.String())
	var get models.Todo
	db.Find(&get, "id = ?", todo.Base.ID)
	if get.Title != "" || get.Status != "" {
		t.Fatalf(`Expected todo to be {Title: "", Status: ""}, got %v`, get)
	}
}
