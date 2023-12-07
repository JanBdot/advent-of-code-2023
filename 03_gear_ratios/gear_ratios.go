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
	sum := 0
	for lineNumber, line := range engineSchematic {
		for characterIndex := 0; characterIndex < len(line); characterIndex++ {
			char := line[characterIndex]
			if string(char) == "*" {
				adjacentNumbers := findAdjacentNumbers(lineNumber, characterIndex)
				length := len(adjacentNumbers)
				if length == 0 || length == 1 {
					continue
				} else {
					gearRatio := 1
					for _, value := range adjacentNumbers {
						gearRatio *= value
					}
					sum += gearRatio
				}
			}

		}
	}

	return sum
}

func findAdjacentNumbers(lineNumber, characterIndex int) []int {
	var adjacentNumbers = make([]int, 0)
	firstRow, lastRow, firstColumn, lastColumn := getCappedDimensions(lineNumber, characterIndex, 1)
	for rowIndex := firstRow; rowIndex <= lastRow; rowIndex++ {
		for columnIndex := firstColumn; columnIndex <= lastColumn; columnIndex++ {
			char := engineSchematic[rowIndex][columnIndex]
			if unicode.IsDigit(rune(char)) {
				number, skipLength := getAdjacentNumber(rowIndex, columnIndex)
				adjacentNumbers = append(adjacentNumbers, number)
				columnIndex += skipLength
			}
		}
	}
	return adjacentNumbers
}

func getAdjacentNumber(rowIndex, columnIndex int) (int, int) {
	numberString := ""
	skipLength := 0
	startIndex := columnIndex
	// Find start of number
	for i := columnIndex; i >= 0; i-- {
		if unicode.IsNumber(rune(engineSchematic[rowIndex][i])) {
			startIndex = i
		} else {
			break
		}
	}
	for i := startIndex; i < len(engineSchematic[rowIndex]); i++ {
		if !unicode.IsNumber(rune(engineSchematic[rowIndex][i])) {
			break
		}
		numberString += string(engineSchematic[rowIndex][i])
		if i > columnIndex {
			skipLength++
		}
	}
	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}
	return number, skipLength
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
	firstRow, lastRow, firstColumn, lastColumn := getCappedDimensions(lineNumber, startIndex, numberLength)

	for lineIndex := firstRow; lineIndex <= lastRow; lineIndex++ {
		for characterIndex := firstColumn; characterIndex <= lastColumn; characterIndex++ {
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

func getCappedDimensions(rowIndex, columnIndex, length int) (firstRow, lastRow, firstColumn, lastColumn int) {
	// Cap the startline if necessary
	if rowIndex == 0 {
		firstRow = 0
	} else {
		firstRow = rowIndex - 1
	}

	// Cap the endline if necessary
	if rowIndex == len(engineSchematic)-1 {
		lastRow = len(engineSchematic) - 1
	} else {
		lastRow = rowIndex + 1
	}

	// Cap the starting character index if necessary
	if columnIndex == 0 {
		firstColumn = 0
	} else {
		firstColumn = columnIndex - 1
	}

	// Cap the ending character index if necessary
	if columnIndex+length == len(engineSchematic[rowIndex]) {
		lastColumn = len(engineSchematic[rowIndex]) - 1
	} else {
		lastColumn = columnIndex + length
	}
	return
}
