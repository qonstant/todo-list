package validators

import (
	"errors"
	"regexp"
	"todo-list/db/sqlc"
)

func ValidateTask(task db.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	if task.ActiveAt.IsZero() {
		return errors.New("activeAt is required")
	}

	// Validate the format of activeAt (YYYY-MM-DD)
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, task.ActiveAt.Format("2006-01-02"))
	if !match {
		return errors.New("activeAt must be in the format YYYY-MM-DD")
	}

	return nil
}
