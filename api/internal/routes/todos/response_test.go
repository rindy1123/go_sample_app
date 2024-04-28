package todos

import (
	"testing"

	"github.com/api/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, expected, got)
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

	assert.Len(t, got.Todos, 2)
}
