package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readMaps(scan *bufio.Scanner) [][][3]int {
	maps := [][][3]int{}
	mapIndex := -1
	inMap := false

	for scan.Scan() {
		if len(scan.Text()) == 0 {
			inMap = false
		} else if !inMap {
			inMap = true
			maps = append(maps, [][3]int{})
			mapIndex++
		} else {
			mapText := scan.Text()
			mapSplit := strings.Split(mapText, " ")
			a, _ := strconv.Atoi(mapSplit[0])
			b, _ := strconv.Atoi(mapSplit[1])
			c, _ := strconv.Atoi(mapSplit[2])
			maps[mapIndex] = append(maps[mapIndex], [3]int{a, b, c})
		}
	}

	return maps
}

func processSeed(seed int, maps [][][3]int) int {
	s := seed
	for _, m := range maps {
		for _, n := range m {
			if s >= n[1] && s <= n[1]+n[2] {
				s = n[0] + s - n[1]
				break
			}
		}
	}

	return s
}

func runDay05(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	seeds := []int{}
	scan.Scan()
	seedsText := scan.Text()
	seedsSplit := strings.Split(seedsText, " ")
	for _, s := range seedsSplit[1:] {
		s, _ := strconv.Atoi(s)
		seeds = append(seeds, s)
	}

	maps := readMaps(scan)
	part1 := -1
	for _, seed := range seeds {
		s := processSeed(seed, maps)
		if part1 == -1 || s < part1 {
			part1 = s
		}
	}

	// Bit of a brute force here, there are ways to shortcut this - but it runs in reasonable time
	// even using just one CPU core.
	part2 := -1
	i := 0
	for i < len(seeds) {
		seed := seeds[i]
		for seed < seeds[i]+seeds[i+1] {
			s := processSeed(seed, maps)
			if part2 == -1 || s < part2 {
				part2 = s
			}
			seed++
		}
		i += 2
	}

	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day05-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay05(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
