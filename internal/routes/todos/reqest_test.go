package todos

import (
	"testing"

	"github.com/api/internal/models"
	"github.com/stretchr/testify/assert"
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
		got := validateTodoStatus(test.status)
		assert.Equal(t, test.expected, got)
	}
}
