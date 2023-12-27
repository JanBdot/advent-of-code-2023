package camelcards

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/JanBdot/advent-of-code-2023/util"
)

var input []string

type round struct {
	cards       []string
	mappedCards map[string]int
	bet         int
}

func CamelCards() (int, int) {
	file, err := util.ReadFile("./07_camel_cards/testinput")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return round1(), round2()
}

func round1() int {
	rounds := make([]round, 0)
	for _, line := range input {
		rounds = append(rounds, mapRound(line))
	}

	sort.Slice(rounds, func(i, j int) bool {
		return CompareRounds(rounds[j], rounds[i], 1)
	})

	sum := 0
	for i := 0; i < len(rounds); i++ {
		sum += rounds[i].bet * (i + 1)
	}
	return sum
}

func round2() int {
	rounds := make([]round, 0)
	for _, line := range input {
		rounds = append(rounds, mapRound(line))
	}

	sort.Slice(rounds, func(i, j int) bool {
		return CompareRounds(rounds[j], rounds[i], 2)
	})

	for _, round := range rounds {
		fmt.Printf("%v\t%v\n", round.cards, round.bet)
	}

	sum := 0
	for i := 0; i < len(rounds); i++ {
		sum += rounds[i].bet * (i + 1)
	}
	return sum
}

func CompareRounds(roundA, roundB round, round int) bool {
	if round == 2 {
		//
	}
	highestSetA := GetHighestSet(roundA.mappedCards, round)
	highestSetB := GetHighestSet(roundB.mappedCards, round)

	if highestSetA > highestSetB {
		return true
	}
	if highestSetB > highestSetA {
		return false
	}
	if highestSetA == 3 && highestSetB == 3 {
		fullHouseA := CheckForFullHouse(roundA.mappedCards)
		fullHouseB := CheckForFullHouse(roundB.mappedCards)
		if fullHouseA && !fullHouseB {
			return true
		}
		if fullHouseB && !fullHouseA {
			return false
		}
	}

	if highestSetA == 2 && highestSetB == 2 {
		fullHouseA := CheckForDoublePair(roundA.mappedCards)
		fullHouseB := CheckForDoublePair(roundB.mappedCards)
		if fullHouseA && !fullHouseB {
			return true
		}
		if fullHouseB && !fullHouseA {
			return false
		}
	}

	// If not returned by now second ordering rule takes effect
	for i := 0; i < len(roundA.cards); i++ {
		cardAValue := GetCardValue(roundA.cards[i], round)
		cardBValue := GetCardValue(roundB.cards[i], round)
		if cardAValue > cardBValue {
			return true
		} else if cardBValue > cardAValue {
			return false
		} else {
			continue
		}
	}

	// exactly the same
	log.Fatal("Edge case similiar hand")
	return true
}

func CheckForDoublePair(cardAmountMap map[string]int) bool {
	amounts := make([]int, len(cardAmountMap))
	for _, amount := range cardAmountMap {
		amounts = append(amounts, amount)
	}
	sort.Ints(amounts)
	if amounts[len(amounts)-1] == 2 && amounts[len(amounts)-2] == 2 {
		return true
	}

	return false
}

func GetCardValue(s string, round int) int {
	if s == "A" {
		return 14
	} else if s == "K" {
		return 13
	} else if s == "Q" {
		return 12
	} else if s == "J" {
		if round == 1 {
			return 11
		}
		if round == 2 {
			return 1
		}
	} else if s == "T" {
		return 10
	} else {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("unexpected Card: %v", err)
		}
		return v
	}
	return -1
}

func CheckForFullHouse(cardAmountMap map[string]int) bool {
	amounts := make([]int, len(cardAmountMap))
	for _, amount := range cardAmountMap {
		amounts = append(amounts, amount)
	}
	sort.Ints(amounts)
	if amounts[len(amounts)-1] == 3 {
		if amounts[len(amounts)-2] == 3 || amounts[len(amounts)-2] == 2 {
			return true
		}
	}

	return false
}

func GetHighestSet(cardAmountMap map[string]int, round int) int {
	jokers := 0
	if round == 2 {
		if jokersTmp, ok := cardAmountMap["J"]; ok {
			jokers = jokersTmp
		}
		delete(cardAmountMap, "J")
	}
	highestAmount := 0
	for _, amount := range cardAmountMap {
		if amount > highestAmount {
			highestAmount = amount
		}
	}
	return highestAmount + jokers
}

func mapRound(line string) round {
	cardsString, betString, _ := strings.Cut(line, " ")
	cards := strings.Split(strings.TrimSpace(cardsString), "")
	mappedCards := MapCards(cards)
	bet, err := strconv.Atoi(betString)
	if err != nil {
		log.Fatalln(err)
	}
	mappedRound := round{
		cards:       cards,
		mappedCards: mappedCards,
		bet:         bet,
	}

	return mappedRound
}

func MapCards(cards []string) map[string]int {
	cardAmountMap := make(map[string]int)
	for i, card := range cards {
		tmpAmount := 0
		for j := i; j < len(cards); j++ {
			if cards[j] == card {
				tmpAmount++
			}
		}
		if _, ok := cardAmountMap[card]; !ok {
			cardAmountMap[card] = tmpAmount
		}
	}
	return cardAmountMap
}
