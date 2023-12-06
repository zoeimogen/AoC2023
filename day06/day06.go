package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func solveRace(t int, d int) int {
	maxTime := (0.0 - float64(t) - math.Sqrt(float64(t*t-4*d))) / -2
	minTime := (0.0 - float64(t) + math.Sqrt(float64(t*t-4*d))) / -2
	solutions := int(math.Floor(maxTime)-math.Ceil(minTime)) + 1
	if maxTime == math.Floor(maxTime) {
		solutions--
	}
	if minTime == math.Ceil(minTime) {
		solutions--
	}
	return solutions
}

func runDay06(inputFile string) (int, int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	times := []int{}
	distances := []int{}

	scan.Scan()
	timesText := scan.Text()
	timesSplit := strings.Split(timesText, " ")
	for _, t := range timesSplit[1:] {
		t, err := strconv.Atoi(t)
		if err == nil {
			times = append(times, t)
		}
	}
	part2TimeText := strings.ReplaceAll(scan.Text(), " ", "")
	part2TimeText = strings.ReplaceAll(part2TimeText, "Time:", "")
	part2Time, _ := strconv.Atoi(part2TimeText)

	scan.Scan()
	distancesText := scan.Text()
	distancesSplit := strings.Split(distancesText, " ")
	for _, d := range distancesSplit[1:] {
		d, err := strconv.Atoi(d)
		if err == nil {
			distances = append(distances, d)
		}
	}
	part2DistanceText := strings.ReplaceAll(scan.Text(), " ", "")
	part2DistanceText = strings.ReplaceAll(part2DistanceText, "Distance:", "")
	part2Distance, _ := strconv.Atoi(part2DistanceText)

	part1 := 1

	for i, t := range times {
		solutions := solveRace(t, distances[i])
		part1 *= solutions
	}

	part2 := solveRace(part2Time, part2Distance)
	return part1, part2
}

func main() {
	var inputFile = flag.String("input", "day06-input.txt", "Problem input file")
	flag.Parse()
	part1, part2 := runDay06(*inputFile)
	fmt.Printf("%d, %d\n", part1, part2)
}
