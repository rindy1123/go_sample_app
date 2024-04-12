package todos

import (
	"testing"

	"github.com/api/internal/models"
)

func TestValidateTodoStatus(t *testing.T) {
	tests := []struct {
		status   models.TodoStatus
		expected bool
	}{
		{"todo", true},
		{"inprogress", true},
		{"done", true},
		{"", true},
		{"invalid", false},
	}

	for _, test := range tests {
		if got := validateTodoStatus(test.status); got != test.expected {
			t.Errorf(
				`Expected ValidateTodoStatus(%v) to be %v, got %v`,
				test.status,
				test.expected,
				got,
			)
		}
	}
}
