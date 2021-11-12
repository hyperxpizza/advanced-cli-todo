package common

import (
	"os"
)

//Checks if file with a provided path exists, if not, returns ErrFileNotFound
func CheckIfFileExists(path string) error {
	_, err := os.Stat(path)
}
