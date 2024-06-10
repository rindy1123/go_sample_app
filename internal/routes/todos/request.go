package todos

import (
	"github.com/api/internal/infra/models"
	"github.com/go-playground/validator/v10"
)

type PostReqBody struct {
	Title  string            `json:"title"  binding:"required"`
	Status models.TodoStatus `json:"status" binding:"required,todostatus"`
}

type UpdateReqBody struct {
	Title  string            `json:"title"`
	Status models.TodoStatus `json:"status" binding:"todostatus"`
}

var todoStatusValidator validator.Func = func(fl validator.FieldLevel) bool {
	status, ok := fl.Field().Interface().(models.TodoStatus)
	if !ok {
		return true
	}
	return validateTodoStatus(status)
}

func validateTodoStatus(status models.TodoStatus) bool {
	switch status {
	case models.TODO, models.INPROGRESS, models.DONE, "":
		return true
	default:
		return false
	}
}
