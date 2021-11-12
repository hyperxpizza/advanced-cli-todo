package common

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

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

//Checks wether provided file is of the right extension
func CheckFileExtension(path, extension string) bool {
	ext := filepath.Ext(path)
	return ext == extension
}
