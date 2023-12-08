package race

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/JanBdot/advent-of-code-2023/util"
)

var fileAsSlice []string

func Race() (int, int) {
	file, err := util.ReadFile("./06_race/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileAsSlice = append(fileAsSlice, scanner.Text())
	}

	return round1(), round2()
}

func round2() int {
	time := getIntFromString(fileAsSlice[0], ":")
	distance := getIntFromString(fileAsSlice[1], ":")

	return len(findBeatingPossiblities(time, distance))
}

func round1() int {
	times := getIntSliceFromString(fileAsSlice[0], ":")
	distances := getIntSliceFromString(fileAsSlice[1], ":")

	sum := 1
	for i := 0; i < len(times); i++ {
		beatingPossibilities := findBeatingPossiblities(times[i], distances[i])
		sum *= len(beatingPossibilities)
	}

	return sum
}

func findBeatingPossiblities(totalTime, minDistance int) []int {
	var possibilities []int
	for i := 1; i < totalTime-1; i++ {
		raceTime := totalTime - i
		speed := i
		if speed*raceTime > minDistance {
			possibilities = append(possibilities, i)
		}
	}
	return possibilities
}

func getIntSliceFromString(s string, sep string) []int {
	_, valuesAsString, _ := strings.Cut(s, sep)
	valuesAsStringSlice := strings.Fields(valuesAsString)
	values := make([]int, len(valuesAsStringSlice))
	for i, valueAsString := range valuesAsStringSlice {
		v, err := strconv.Atoi(valueAsString)
		if err != nil {
			panic(err)
		}
		values[i] = v
	}
	return values
}

func getIntFromString(s string, sep string) int {
	_, valueAsString, _ := strings.Cut(s, sep)
	valueAsString = strings.ReplaceAll(valueAsString, " ", "")
	v, err := strconv.Atoi(valueAsString)
	if err != nil {
		panic(err)
	}
	return v
}
