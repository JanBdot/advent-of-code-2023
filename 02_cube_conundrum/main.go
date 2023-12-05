package main

import (
	"adc/util"
	"bufio"
	"fmt"
)

type Round struct {
	RoundId int
	// Each round has a list of draws with a map of the amount of colors
	Draw []map[string]int
}

func main() {
	file, err := util.ReadFile("testinput")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rounds []Round
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		rounds = append(rounds, MapInputToRound(scanner.Text()))
	}
}

func MapInputToRound(s string) Round {
	panic("unimplemented")
}
