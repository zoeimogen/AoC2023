package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func runDay01Part1(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	re := regexp.MustCompile(`[0-9]`)
	total := 0

	for scan.Scan() {
		d := re.FindAll([]byte(scan.Text()), -1)
		digits := string(d[0]) + string(d[len(d)-1])

		i, _ := strconv.Atoi(digits)

		if err != nil {
			println("Something went wrong")
			os.Exit(-1)
		}
		total += i
	}

	return total
}

func readDigit(z string) string {
	if z == "one" {
		return "1"
	}
	if z == "two" {
		return "2"
	}
	if z == "three" {
		return "3"
	}
	if z == "four" {
		return "4"
	}
	if z == "five" {
		return "5"
	}
	if z == "six" {
		return "6"
	}
	if z == "seven" {
		return "7"
	}
	if z == "eight" {
		return "8"
	}
	if z == "nine" {
		return "9"
	}
	return z
}

func runDay01Part2(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	re := regexp.MustCompile(`[0-9]|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)
	total := 0

	for scan.Scan() {
		text := []byte(scan.Text())

		d := re.FindAll(text, -1)
		a := readDigit(string(d[0]))
		b := readDigit(string(d[len(d)-1]))
		digits := a + b

		i, _ := strconv.Atoi(digits)

		if err != nil {
			println("Something went wrong")
			os.Exit(-1)
		}
		total += i
	}

	return total
}

func main() {
	var inputFile = flag.String("input", "day01-input.txt", "Problem input file")
	flag.Parse()
	part1 := runDay01Part1(*inputFile)
	part2 := runDay01Part2(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
