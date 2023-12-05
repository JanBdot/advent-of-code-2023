package gear_ratios

import (
	"bufio"
	"os"
	"strconv"
	"unicode"

	"github.com/JanBdot/advent-of-code-2023/util"
)

var engineSchematic []string

func GearRatios() (int, int) {
	file, err := util.ReadFile("./03_gear_ratios/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		engineSchematic = append(engineSchematic, scanner.Text())
	}

	return round1(file), round2(file)
}

func round2(file *os.File) int {
	return 0
}

func round1(file *os.File) int {
	sum := 0
	for lineNumber, line := range engineSchematic {
		for characterIndex := 0; characterIndex < len(line); {
			if unicode.IsDigit(rune(line[characterIndex])) {
				numberValue, numberLength := getNumber(line, characterIndex)
				if adjacentToSymbol(lineNumber, characterIndex, numberLength) {
					sum += numberValue
				}
				characterIndex += numberLength
			} else {
				characterIndex++
			}

		}
	}

	return sum
}

func adjacentToSymbol(lineNumber, startIndex, numberLength int) bool {
	var startLine, endLine, startChar, endChar int

	// Cap the startline if necessary
	if lineNumber == 0 {
		startLine = 0
	} else {
		startLine = lineNumber - 1
	}

	// Cap the endline if necessary
	if lineNumber == len(engineSchematic)-1 {
		endLine = len(engineSchematic) - 1
	} else {
		endLine = lineNumber + 1
	}

	// Cap the starting character index if necessary
	if startIndex == 0 {
		startChar = 0
	} else {
		startChar = startIndex - 1
	}

	// Cap the ending character index if necessary
	if startIndex+numberLength == len(engineSchematic[lineNumber]) {
		endChar = len(engineSchematic[lineNumber]) - 1
	} else {
		endChar = startIndex + numberLength
	}

	for lineIndex := startLine; lineIndex <= endLine; lineIndex++ {
		for characterIndex := startChar; characterIndex <= endChar; characterIndex++ {
			char := engineSchematic[lineIndex][characterIndex]
			if string(char) == "." {
				continue
			} else if unicode.IsSymbol(rune(char)) || unicode.IsPunct(rune(char)) {
				return true
			}
		}
	}

	return false
}

func getNumber(line string, startIndex int) (int, int) {
	var valueString string
	length := 0
	for i := startIndex; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			valueString += string(line[i])
			length++
		} else {
			break
		}
	}
	value, err := strconv.Atoi(valueString)
	if err != nil {
		panic(err)
	}
	return value, length
}
