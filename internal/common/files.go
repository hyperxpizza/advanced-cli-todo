package common

import (
	"errors"
	"os"

	"github.com/hyperxpizza/advanced-cli-todo/internal/customErrors"
)

//Checks if file with a provided path exists, if not, returns ErrFileNotFound
func CheckIfFileExists(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return customErrors.Wrap(customErrors.ErrFileNotFound)
	}

	return nil
}
