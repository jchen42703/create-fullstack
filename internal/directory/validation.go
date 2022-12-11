package directory

import (
	"errors"
	"os"
)

// exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
	if _, err := os.Stat(path); err == nil {
		// path/to/whatever exists
		return true, nil

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		return false, nil
	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		return false, err
	}
}
