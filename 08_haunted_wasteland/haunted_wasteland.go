package haunted_wasteland

import (
	"os"
	"regexp"
	"strings"
)

func parseInput(lines []string) (map[string][2]string, string) {
	instructions := lines[0]
	path := lines[1]
	re := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

	moveMap := make(map[string][2]string)
	for _, m := range re.FindAllStringSubmatch(path, -1) {
		moveMap[strings.TrimSpace(m[1])] = [2]string{strings.TrimSpace(m[2]), strings.TrimSpace(m[3])}
	}

	return moveMap, instructions
}

func Round1(input []byte) int {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	moveMap, instructions := parseInput(lines)
	start := "AAA"
	end := "ZZZ"
	steps := 0

	for start != end {
		for _, instruction := range instructions {
			if instruction == 'L' {
				start = moveMap[start][0]
			} else {
				start = moveMap[start][1]
			}
			steps++
			if start == end {
				break
			}
		}
	}

	return steps
}

func Round2() int {
	// input, err := readInput()
	// if err != nil {
	// 	return 0, err
	// }

	// TODO: Implement the logic for round 2

	return 0
}

func HauntedWasteland() (int, int) {
	input, err := os.ReadFile("08_haunted_wasteland/input")
	if err != nil {
		panic(err)
	}

	round1 := Round1(input)
	round2 := Round2()
	return round1, round2
}
