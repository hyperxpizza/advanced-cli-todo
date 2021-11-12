package common

import (
	"errors"
	"io/ioutil"
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

//Reads whole file into a byte array
func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, customErrors.Wrap(customErrors.ErrFileNotFound)
		}
		return nil, err
	}

	return data, err
}
