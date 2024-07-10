package validators

import (
	"errors"
	"todo-list/internal/models"
)

func ValidateTask(task models.Task) error {
	if len(task.Title) == 0 || len(task.Title) > 200 {
		return errors.New("title is required and must not exceed 200 characters")
	}

	if task.ActiveAt.IsZero() {
		return errors.New("ActiveAt date is required and must be valid")
	}

	return nil
}
