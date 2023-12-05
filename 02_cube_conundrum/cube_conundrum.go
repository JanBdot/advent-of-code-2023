package cubeconundrum

import (
	"bufio"
	"strconv"
	"strings"

	trebuchet "github.com/JanBdot/advent-of-code-2023/01_trebuchet"
	"github.com/JanBdot/advent-of-code-2023/util"
)

var conditionMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Round struct {
	RoundId int
	// Each round has a list of draws with a map of the amount of colors
	Draws []map[string]int
}

func CubeConundrum() int {
	file, err := util.ReadFile("./02_cube_conundrum/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rounds []Round
	for scanner.Scan() {
		rounds = append(rounds, MapInputToRound(scanner.Text()))
	}
	sum := 0
	for _, round := range rounds {
		if DrawsMeetCondition(round.Draws) {
			sum += round.RoundId
		}
	}
	return sum
}

func DrawsMeetCondition(draws []map[string]int) bool {
	for _, draw := range draws {
		for color, amount := range draw {
			if conditionMap[color] < amount {
				return false
			}
		}
	}
	return true
}

func MapInputToRound(s string) Round {
	gameInfoString, drawsString, _ := strings.Cut(s, ":")
	id := ExtractGameId(&gameInfoString)
	draws := GetDrawsMaps(drawsString)
	return Round{RoundId: id, Draws: draws}
}

func GetDrawsMaps(drawsString string) []map[string](int) {
	var draws []map[string](int)
	drawsSlice := strings.Split(drawsString, ";")
	for _, drawString := range drawsSlice {
		draws = append(draws, MapDraw(drawString))
	}
	return draws
}

func MapDraw(drawString string) map[string](int) {
	drawMap := make(map[string]int)
	drawEntries := strings.Split(drawString, ",")
	for _, drawEntry := range drawEntries {
		amountString, color, _ := strings.Cut(strings.TrimSpace(drawEntry), " ")
		amount, err := strconv.Atoi(amountString)
		if err != nil {
			panic(err)
		}
		drawMap[color] = amount
	}
	return drawMap
}

func ExtractGameId(s *string) int {
	trebuchet.ExtractDigits(s)
	id, err := strconv.Atoi(*s)
	if err != nil {
		panic(err)
	}
	return id
}
