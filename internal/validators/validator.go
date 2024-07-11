package validators

import (
	"errors"
	"todo-list/db/sqlc"
)

func ValidateTask(task db.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	if task.ActiveAt.IsZero() {
		return errors.New("activeAt is required and must be a valid date")
	}
	return nil
}
