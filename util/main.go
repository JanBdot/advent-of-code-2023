package util

import "os"

// ReadFile returns a pointer to an opened file and any error encountered
func ReadFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
