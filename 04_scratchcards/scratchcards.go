package scratchcards

import (
	"bufio"
	"math"
	"strings"

	"github.com/JanBdot/advent-of-code-2023/util"
)

var scratchcards []string
var scoresPerLine []int
var cardsAmount []int

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
	for i, score := range scoresPerLine {
		if score > 0 {
			for j := 0; j < score; j++ {
				// cap to end of scoresPerLine
				index := int(math.Min(float64(i+j+1), float64(len(scoresPerLine)-1)))
				// Add the copies of the current line card to the amount
				cardsAmount[index] += cardsAmount[i]
			}
		}
	}
	sum := 0
	for _, amount := range cardsAmount {
		sum += amount
	}
	return sum
}

func round1() int {
	sum := 0
	for _, line := range scratchcards {
		drawnNumbers, winningNumbers := ExtractNumbersFromScratchcard(line)
		correctNumbers := FindWinningNumbers(drawnNumbers, winningNumbers)
		if len(correctNumbers) > 0 {
			sum += int(math.Pow(2, float64(len(correctNumbers)-1)))
		}
		scoresPerLine = append(scoresPerLine, len(correctNumbers))
		cardsAmount = append(cardsAmount, 1)
	}
	return sum
}

func ExtractNumbersFromScratchcard(line string) ([]int, []int) {
	_, numbers, _ := strings.Cut(line, ":")
	drawnNumbersString, winningNumbersString, _ := strings.Cut(numbers, "|")
	drawnNumbers := util.ConvertStringToIntSlice(strings.TrimSpace(drawnNumbersString))
	winningNumbers := util.ConvertStringToIntSlice(strings.TrimSpace(winningNumbersString))
	return drawnNumbers, winningNumbers
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
