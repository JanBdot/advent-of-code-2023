package main

import (
	"adc/util"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func ExtractDigits(s *string) {
	var digits string
	for _, c := range *s {
		if unicode.IsDigit(c) {
			digits += string(c)
		}
	}
	*s = digits
}

func ExtractFirstAndLastDigit(s *string) {
	if len(*s) == 0 {
		*s = "00"
	}
	var twoDigitNumber string
	if len(*s) == 1 {
		twoDigitNumber := fmt.Sprintf("%s%s", *s, *s)
		*s = twoDigitNumber
	}
	stringValue := *s
	first := string(stringValue[0])
	second := string(stringValue[len(*s)-1])
	twoDigitNumber = fmt.Sprintf("%s%s", first, second)
	*s = twoDigitNumber
}

func ConvertWrittenDigits(s *string) {
	writtenDigits := map[string]string{
		"ONE": "1", "TWO": "2", "THREE": "3", "FOUR": "4",
		"FIVE": "5", "SIX": "6", "SEVEN": "7", "EIGHT": "8", "NINE": "9",
	}

	var result strings.Builder

	// Iterate over each character and replace first letter with digit to catch overlaps
	for i := 0; i < len(*s); i++ {
		matched := false
		for word, digit := range writtenDigits {
			if strings.HasPrefix((*s)[i:], word) {
				result.WriteString(digit)
				matched = true
				break
			}
		}
		if !matched {
			result.WriteByte((*s)[i])
		}
	}

	*s = result.String()
}

func main() {
	file, err := util.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int = 0
	for scanner.Scan() {
		// loops each line
		var text = strings.ToUpper(scanner.Text())

		ConvertWrittenDigits(&text)
		ExtractDigits(&text)
		ExtractFirstAndLastDigit(&text)
		add, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		sum += add
	}

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	log.Println(sum)
}
