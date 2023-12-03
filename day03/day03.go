package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	input     = [][]byte{}
	total     = 0
	gearTotal = 0
)

func checkNumber(x int, y int) int {
	// Check a given square for a number
	if x < 0 || y < 0 || x >= len(input[0]) || y >= len(input) {
		// Input values out of range
		return 0
	}

	if input[y][x] < '0' || input[y][x] > '9' {
		// Not a number
		return 0
	}

	// Found a number, scan forward/back to find start and end.
	start := x
	for start > 0 && input[y][start-1] >= '0' && input[y][start-1] <= '9' {
		start--
	}
	end := x
	for end+1 < len(input[y]) && input[y][end+1] >= '0' && input[y][end+1] <= '9' {
		end++
	}

	value, _ := strconv.Atoi(string(input[y][start : end+1]))
	// fmt.Printf("Found %d at %d,%d\n", value, x, y)
	total += value

	// Blank out the found value so it can't be reused
	i := start
	for i <= end {
		input[y][i] = '.'
		i++
	}

	return value
}

func processSymbol(x int, y int) {
	// Check all round a found symbol for numbers
	results := []int{}

	result := checkNumber(x-1, y-1)
	if result > 0 {
		results = append(results, result)
	}
	result = checkNumber(x, y-1)
	if result > 0 {
		results = append(results, result)
	}
	result = checkNumber(x+1, y-1)
	if result > 0 {
		results = append(results, result)
	}
	result = checkNumber(x-1, y)
	if result > 0 {
		results = append(results, result)
	}
	result = checkNumber(x, y)
	if result > 0 {
		results = append(results, result)
	}
	result = checkNumber(x+1, y)
	if result > 0 {
		results = append(results, result)
	}
	result = checkNumber(x-1, y+1)
	if result > 0 {
		results = append(results, result)
	}
	result = checkNumber(x, y+1)
	if result > 0 {
		results = append(results, result)
	}
	result = checkNumber(x+1, y+1)
	if result > 0 {
		results = append(results, result)
	}

	if input[y][x] == '*' && len(results) == 2 {
		gearTotal += results[0] * results[1]
	}
}

func runDay03(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		// Load data
		input = append(input, []byte(scan.Text()))
	}

	for y := range input {
		for x, c := range input[y] {
			if c != '.' && (c < '0' || c > '9') {
				processSymbol(x, y)
			}
		}
	}

	return total, gearTotal
}

func main() {
	var inputFile = flag.String("input", "day03-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay03(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
