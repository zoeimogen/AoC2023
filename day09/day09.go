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

func extrapolate(numbers []int) int {
	diffs := []int{}
	allZero := true

	for i, _ := range numbers[1:] {
		diff := numbers[i+1] - numbers[i]
		if diff != 0 {
			allZero = false
		}

		diffs = append(diffs, diff)
	}

	if allZero {
		return numbers[0]
	}

	result := numbers[len(numbers)-1] + extrapolate(diffs)
	return result
}

func runDay09(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	for scan.Scan() {
		numbersText := strings.Split(scan.Text(), " ")
		numbers := []int{}
		for _, t := range numbersText {
			i, _ := strconv.Atoi(t)
			numbers = append(numbers, i)
		}
		part1 += extrapolate(numbers)
		slices.Reverse(numbers)
		part2 += extrapolate(numbers)
	}

	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day09-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay09(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
