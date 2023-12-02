package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Max Balls count for Part 1
	maxBalls = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func runDay02(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	re := regexp.MustCompile(`(Game [0-9]+)|([0-9]+ ((blue)|(green)|(red))*)`)
	total := 0
	totalPower := 0

	for scan.Scan() {
		text := []byte(scan.Text())
		d := re.FindAll(text, -1)

		game, _ := strconv.Atoi(strings.Split(string(d[0]), " ")[1])

		minBalls := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		ok := true
		for _, balls := range d[1:] {
			split := strings.Split(string(balls), " ")
			ballCount, _ := strconv.Atoi(split[0])

			// Part 1 solution
			if ballCount > maxBalls[split[1]] {
				ok = false
			}

			// Part 2 solution
			if ballCount > minBalls[split[1]] {
				minBalls[split[1]] = ballCount
			}
		}

		// Part 1 solution
		if ok {
			total += game
		}

		// Part 2 solution
		totalPower += minBalls["red"] * minBalls["green"] * minBalls["blue"]
	}

	return total, totalPower
}

func main() {
	var inputFile = flag.String("input", "day02-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay02(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
