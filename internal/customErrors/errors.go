package customErrors

import "errors"

const (
	ErrFileNotFound = "File was not found"
)

//Wraps a predefiined message and returns it as a new error
func Wrap(e string) error {
	return errors.New(e)
}
