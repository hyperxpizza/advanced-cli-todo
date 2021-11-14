package customErrors

import "errors"

const (
	ErrFileNotFound = "File was not found"
	ErrTaskNotValid = "The task is not valid"
)

//Wraps a predefiined message and returns it as a new error
func Wrap(e string) error {
	return errors.New(e)
}
