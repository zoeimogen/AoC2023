package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func processCard(card string) (int, int) {
	score := 0
	wins := 0

	input := strings.Split(card, ": ")
	numbers := strings.Split(input[1], " | ")
	winners := []int{}
	for _, n := range strings.Split(numbers[0], " ") {
		number, err := strconv.Atoi(n)
		if err == nil {
			winners = append(winners, number)
		}
	}
	for _, n := range strings.Split(numbers[1], " ") {
		number, err := strconv.Atoi(n)
		if err == nil && slices.Contains(winners, number) {
			wins += 1
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score, wins
}

func processCardStack(cards []string, start int, count int) (int, int) {
	totalScore := 0
	totalCards := 0

	for i, card := range cards[start : start+count] {
		score, wins := processCard(card)
		totalScore += score
		totalCards += 1

		if wins > 0 {
			_, wins := processCardStack(cards, start+i+1, wins)
			totalCards += wins
		}
	}

	return totalScore, totalCards
}

func runDay04(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	cards := []string{}

	for scan.Scan() {
		// Load data
		cards = append(cards, scan.Text())
	}

	return processCardStack(cards, 0, len(cards))
}

func main() {
	var inputFile = flag.String("input", "day04-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay04(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
