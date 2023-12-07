package scratchcards

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/JanBdot/advent-of-code-2023/util"
)

var scratchcards []string

func Scratchcards() (int, int) {
	file, err := util.ReadFile("./04_scratchcards/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scratchcards = append(scratchcards, scanner.Text())
	}

	return round1(), round2()
}

func round2() int {
	return 0
}

func round1() int {
	sum := 0
	for _, line := range scratchcards {
		drawnNumbers, winningNumbers := ExtractNumbersFromScratchcard(line)
		correctNumbers := FindWinningNumbers(drawnNumbers, winningNumbers)
		if len(correctNumbers) > 0 {
			sum += int(math.Pow(2, float64(len(correctNumbers)-1)))
		}
	}
	return sum
}

func ExtractNumbersFromScratchcard(line string) ([]int, []int) {
	_, numbers, _ := strings.Cut(line, ":")
	drawnNumbersString, winningNumbersString, _ := strings.Cut(numbers, "|")
	drawnNumbers := ConvertStringToIntSlice(strings.TrimSpace(drawnNumbersString))
	winningNumbers := ConvertStringToIntSlice(strings.TrimSpace(winningNumbersString))
	return drawnNumbers, winningNumbers
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

func FindWinningNumbers(drawnNumbers, winningNumbers []int) []int {
	var numbers = make([]int, 0)
	for _, drawnNumber := range drawnNumbers {
		for _, winningNumber := range winningNumbers {
			if drawnNumber == winningNumber {
				numbers = append(numbers, drawnNumber)
			}
		}
	}
	return numbers
}
