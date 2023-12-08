package util

import (
	"os"
	"strconv"
	"strings"
)

// ReadFile returns a pointer to an opened file and any error encountered
func ReadFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ConvertStringToIntSlice(numbersString string) []int {
	var intSlice = make([]int, 0)
	stringSlice := strings.Fields(numbersString)
	for _, numberString := range stringSlice {
		number, err := strconv.Atoi(string(numberString))
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, number)
	}
	return intSlice
}
