package customErrors

import "errors"

const (
	ErrFileNotFound         = "File was not found"
	ErrTaskTitleNotValid    = "The task title is not valid"
	ErrTaskPriorityNotValid = "The task priority is not valid"
)

//Wraps a predefiined message and returns it as a new error
func Wrap(e string) error {
	return errors.New(e)
}
