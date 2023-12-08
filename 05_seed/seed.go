package seed

import (
	"bufio"
	"slices"
	"strings"

	"github.com/JanBdot/advent-of-code-2023/util"
)

var fileAsSlice []string
var seedList []int

var maps map[string][][]int

func Seed() (int, int) {
	file, err := util.ReadFile("./05_seed/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileAsSlice = append(fileAsSlice, scanner.Text())
	}

	// order matters here!
	initSeedList()
	initMaps()

	return round1(), round2()
}

func initMaps() {
	maps = make(map[string][][]int)
	for lineNumber := 0; lineNumber < len(fileAsSlice); lineNumber++ {
		mapName, _, foundMap := strings.Cut(fileAsSlice[lineNumber], "map:")
		if foundMap {
			mapLines := make([][]int, 0)
			for foundMap {
				lineNumber++
				if lineNumber >= len(fileAsSlice) {
					break
				}
				line := fileAsSlice[lineNumber]
				if line == "" {
					foundMap = false
				} else {
					mapLines = append(mapLines, util.ConvertStringToIntSlice(line))
				}
			}
			mapName = strings.TrimSpace(mapName)
			maps[mapName] = mapLines
		}
	}
}

func initSeedList() {
	_, seedsString, _ := strings.Cut(fileAsSlice[0], ":")
	fileAsSlice = fileAsSlice[2:]
	seedList = util.ConvertStringToIntSlice(seedsString)
}

func round2() int {

	locations := make([]int, 0)
	for i := 0; i < len(seedList); i += 2 {
		startSeed := seedList[i]
		seedRange := seedList[i+1]
		tmpLocations := make([]int, 0)
		// bruteforce attempt (takes a long time)
		for seed := startSeed; seed < startSeed+seedRange; seed++ {
			tmpDest := seed
			getDestination("seed-to-soil", &tmpDest)
			getDestination("soil-to-fertilizer", &tmpDest)
			getDestination("fertilizer-to-water", &tmpDest)
			getDestination("water-to-light", &tmpDest)
			getDestination("light-to-temperature", &tmpDest)
			getDestination("temperature-to-humidity", &tmpDest)
			getDestination("humidity-to-location", &tmpDest)
			tmpLocations = append(tmpLocations, tmpDest)
		}
		locations = append(locations, slices.Min(tmpLocations))
	}

	return slices.Min(locations)
}

func round1() int {
	locations := make([]int, 0)
	for _, seed := range seedList {
		tmpDest := seed
		getDestination("seed-to-soil", &tmpDest)
		getDestination("soil-to-fertilizer", &tmpDest)
		getDestination("fertilizer-to-water", &tmpDest)
		getDestination("water-to-light", &tmpDest)
		getDestination("light-to-temperature", &tmpDest)
		getDestination("temperature-to-humidity", &tmpDest)
		getDestination("humidity-to-location", &tmpDest)
		locations = append(locations, tmpDest)
	}

	return slices.Min(locations)
}

func getDestination(s string, tmpSource *int) {
	finalDestination := *tmpSource
	for _, mapLines := range maps[s] {
		destination := mapLines[0]
		source := mapLines[1]
		rangeLength := mapLines[2]

		if *tmpSource >= source && *tmpSource < source+rangeLength {
			finalDestination = *tmpSource + destination - source
		}
	}
	*tmpSource = finalDestination
}
