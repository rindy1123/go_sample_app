package models_test

import (
	"testing"

	"github.com/api/internal/infra/models"
	"github.com/api/test"
	"github.com/api/test/factory"
	"github.com/stretchr/testify/assert"
)

func TestTodo_Create(t *testing.T) {
	db := test.SetupTestDB(t)
	todo := models.Todo{Title: "test", Status: "status"}
	id, err := todo.Create(db)
	assert.NoError(t, err)

	var got models.Todo
	db.Find(&got, "id = ?", id)
	assert.Equal(t, todo.Title, got.Title)
	assert.Equal(t, todo.Status, got.Status)
}

func TestTodo_Get(t *testing.T) {
	db := test.SetupTestDB(t)
	todo := factory.Factory{}.CreateTodo(db)
	got := models.Todo{}.Get(db, todo.Base.ID.String())
	assert.Equal(t, todo.Title, got.Title)
	assert.Equal(t, todo.Status, got.Status)
}

func TestTodo_List(t *testing.T) {
	db := test.SetupTestDB(t)
	factory.Factory{}.CreateTodo(db)
	factory.Factory{}.CreateTodo(db)
	todos := models.Todo{}.List(db)
	assert.Len(t, todos, 2)
}

func TestTodo_Update(t *testing.T) {
	db := test.SetupTestDB(t)
	todo := factory.Factory{}.CreateTodo(db)
	got := models.Todo{Title: "updated"}.Update(db, todo.Base.ID.String())
	assert.Equal(t, "updated", got.Title)
	assert.Equal(t, todo.Status, got.Status)
}

func TestTodo_HardDelete(t *testing.T) {
	db := test.SetupTestDB(t)
	todo := factory.Factory{}.CreateTodo(db)
	models.Todo{}.HardDelete(db, todo.Base.ID.String())
	var got models.Todo
	db.Find(&got, "id = ?", todo.Base.ID)
	assert.Equal(t, models.Todo{}, got)
}
