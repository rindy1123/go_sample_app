package todos

import (
	"testing"

	"github.com/api/internal/models"
	"github.com/google/uuid"
)

func TestToSingular(t *testing.T) {
	id := uuid.New()
	todo := models.Todo{
		Base:   models.Base{ID: id},
		Title:  "Test",
		Status: models.TODO,
	}

	got := toSingular(todo)
	expected := GetTodo{
		ID:     id,
		Title:  "Test",
		Status: models.TODO,
	}

	if got != expected {
		t.Errorf("Expected toSingular(%v) to be %v, got %v", todo, expected, got)
	}
}

func TestToMultiple(t *testing.T) {
	id := uuid.New()
	todos := []models.Todo{
		{
			Base:   models.Base{ID: id},
			Title:  "Test",
			Status: models.TODO,
		},
		{
			Base:   models.Base{ID: id},
			Title:  "Test",
			Status: models.TODO,
		},
	}

	got := toMultiple(todos)

	if len(got.Todos) != 2 {
		t.Errorf("Expected 2 elements, got %v", len(got.Todos))
	}
}
