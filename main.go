package main

import (
	"fmt"

	// trebuchet "github.com/JanBdot/advent-of-code-2023/01_trebuchet"
	// cubeconundrum "github.com/JanBdot/advent-of-code-2023/02_cube_conundrum"
	// gear_ratios "github.com/JanBdot/advent-of-code-2023/03_gear_ratios"
	// scratchcards "github.com/JanBdot/advent-of-code-2023/04_scratchcards"
	// seed "github.com/JanBdot/advent-of-code-2023/05_seed"
	race "github.com/JanBdot/advent-of-code-2023/06_race"
)

func main() {
	var round1, round2 int
	fmt.Println("Advent of Code 2023 by Jan Baer")
	// round1, round2 = trebuchet.Trebuchet()
	// fmt.Printf("Day 01:\n\tround 1: %v\n\tround 2: %v\n", round1, round2)
	// round1, round2 = cubeconundrum.CubeConundrum()
	// fmt.Printf("Day 02:\n\tround 1: %v\n\tround 2: %v\n", round1, round2)
	// round1, round2 = gear_ratios.GearRatios()
	// fmt.Printf("Day 03:\n\tround 1: %v\n\tround 2: %v\n", round1, round2)
	// round1, round2 = scratchcards.Scratchcards()
	// fmt.Printf("Day 04:\n\tround 1: %v\n\tround 2: %v\n", round1, round2)
	// round1, round2 = seed.Seed()
	// fmt.Printf("Day 05:\n\tround 1: %v\n\tround 2: %v\n", round1, round2)
	round1, round2 = race.Race()
	fmt.Printf("Day 06:\n\tround 1: %v\n\tround 2: %v\n", round1, round2)
}
