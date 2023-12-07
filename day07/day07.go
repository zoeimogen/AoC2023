package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards      []byte
	bid        int
	scorePart1 int
	scorePart2 int
}

func scoreCard(card byte, jokersWild bool) int {
	if card == 'A' {
		return 14
	} else if card == 'K' {
		return 13
	} else if card == 'Q' {
		return 12
	} else if card == 'J' {
		if jokersWild {
			return 1
		} else {
			return 11
		}
	} else if card == 'T' {
		return 10
	}
	return int(card - '0')
}

func scoreCards(cards []byte, jokersWild bool) int {
	// Base card score is 4 bits per card.
	score := 0
	for _, card := range cards {
		score *= 16
		score += scoreCard(card, jokersWild)
	}
	return score
}

func compareCards(a, b byte, jokersWild bool) bool {
	if jokersWild && (a == 'J' || b == 'J') {
		return true
	}
	return a == b
}

func getHandScore(cards []byte, jokersWild bool) int {
	// Total hand score is the base score (4 bits per card, max 978670) plus
	// some multiple of 1,000,000 representing the value of the hand type.

	score := scoreCards(cards, jokersWild)

	// Count matches. Dirty but works, should be a single recursive function really.
	matches := 0
	for _, c := range cards[1:] {
		if compareCards(cards[0], c, jokersWild) {
			matches++
		}
	}
	for _, c := range cards[2:] {
		if compareCards(cards[1], c, jokersWild) {
			matches++
		}
	}
	for _, c := range cards[3:] {
		if compareCards(cards[2], c, jokersWild) {
			matches++
		}
	}
	if compareCards(cards[3], cards[4], jokersWild) {
		matches++
	}

	// If jokers are wild, we need to know how many there are - becase that distorts the
	// expected number of matches for some hands.
	jokerCount := 0

	if jokersWild {
		for _, c := range cards {
			if c == 'J' {
				jokerCount++
			}
		}
	}

	// Five of a kind: Score 6000000 + cards vale
	if matches == 10 {
		return 6000000 + score
	}

	// Four of a kind: Score 5000000 + cards value
	if matches == 6+jokerCount {
		return 5000000 + score
	}

	// Full House etc
	if matches == 4+2*jokerCount {
		return 4000000 + score
	}

	// Three of a kind
	if matches == 3+2*jokerCount {
		return 3000000 + score
	}

	// Two pair
	if matches == 2 {
		return 2000000 + score
	}

	// One pair
	if matches == 1+3*jokerCount {
		return 1000000 + score
	}

	// High card
	return score
}

func runDay07(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	hands := []Hand{}

	for scan.Scan() {
		var cards []byte
		hand := strings.Split(scan.Text(), " ")
		cards = []byte(hand[0])
		bid, _ := strconv.Atoi(hand[1])
		hands = append(hands, Hand{cards, bid,
			getHandScore(cards, false),
			getHandScore(cards, true)})
	}

	slices.SortFunc(hands, func(a, b Hand) int { return cmp.Compare(a.scorePart1, b.scorePart1) })
	part1 := 0

	for rank, hand := range hands {
		part1 += (rank + 1) * hand.bid
	}

	println("")
	slices.SortFunc(hands, func(a, b Hand) int { return cmp.Compare(a.scorePart2, b.scorePart2) })
	part2 := 0

	for rank, hand := range hands {
		part2 += (rank + 1) * hand.bid
	}

	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day07-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay07(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
