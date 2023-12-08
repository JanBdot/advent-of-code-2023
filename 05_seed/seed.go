package seed

import (
	"bufio"

	"github.com/JanBdot/advent-of-code-2023/util"
)

func Seed() (int, int) {
	file, err := util.ReadFile("./05_seed/testinput")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}

	return round1(), round2()
}

func round2() int {
	return 0
}

func round1() int {
	return 0
}
