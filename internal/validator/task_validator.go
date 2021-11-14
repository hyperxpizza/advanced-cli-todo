package validator

import (
	"regexp"

	"github.com/hyperxpizza/advanced-cli-todo/internal/customErrors"
	"github.com/hyperxpizza/advanced-cli-todo/internal/models"
)

//Validates a new task
func ValidateNewTask(t models.Task) error {
	var isTitleValid = regexp.MustCompile(`^(\S|\S.{0,100}\S)$`).MatchString
	if !isTitleValid(t.Title) {
		return customErrors.Wrap(customErrors.ErrTaskTitleNotValid)
	}

	if t.Priority > 10 {
		return customErrors.Wrap(customErrors.ErrTaskPriorityNotValid)
	}

	return nil
}
