package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Node struct {
	left  [3]byte
	right [3]byte
}

type Nodes map[[3]byte]Node

func loadData(inputFile string) (string, Nodes) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	nodes := Nodes{}
	scan.Scan()
	path := scan.Text()
	scan.Scan()

	for scan.Scan() {
		var thisNode []byte
		var left []byte
		var right []byte
		var node Node
		var thisNodeBytes [3]byte

		n, err := fmt.Sscanf(scan.Text(), "%3s = (%3s, %3s)", &thisNode, &left, &right)
		if n != 3 || err != nil {
			log.Fatalf("Can't parse \"%s\": %v\n", scan.Text(), err)
		}
		copy(node.left[:], left)
		copy(node.right[:], right)
		copy(thisNodeBytes[:], thisNode)
		nodes[thisNodeBytes] = node
	}

	return path, nodes
}

func runDay08Part1(path string, nodes Nodes) int {
	part1 := 0
	currentNode := [3]byte{'A', 'A', 'A'}

	for currentNode != [3]byte{'Z', 'Z', 'Z'} {
		for n, lr := range path {
			part1++
			if lr == 'L' {
				currentNode = nodes[currentNode].left
			} else if lr == 'R' {
				currentNode = nodes[currentNode].right
			} else {
				log.Fatalf("Unknown path element %d: '%c'", n, lr)
			}
		}
	}

	return part1
}

func greatestCommonDivisor(x, y int) int {
	// https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}

func leastCommonMultiple(ints []int) int {
	// https://en.wikipedia.org/wiki/Least_common_multiple#Using_the_greatest_common_divisor
	result := ints[0] * ints[1] / greatestCommonDivisor(ints[0], ints[1])

	for _, i := range ints[2:] {
		result = leastCommonMultiple([]int{i, result})
	}

	return result
}

func runDay08Part2(path string, nodes Nodes) int {
	count := 0
	currentNodes := [][3]byte{}
	finishedNodes := []int{}

	for n, _ := range nodes {
		if n[2] == 'A' {
			currentNodes = append(currentNodes, n)
			finishedNodes = append(finishedNodes, 0)
		}
	}

	// Don't try to solve it by brute force - for each starting node, just
	// iterate until it finishes, then find the lowest common multiple, being
	// the first point all of the nodes finish simultaneously.
	for {
		for n, lr := range path {
			count++

			if lr == 'L' {
				for i, node := range currentNodes {
					if finishedNodes[i] == 0 {
						currentNodes[i] = nodes[node].left
					}
				}
			} else if lr == 'R' {
				for i, node := range currentNodes {
					if finishedNodes[i] == 0 {
						currentNodes[i] = nodes[node].right
					}
				}
			} else {
				log.Fatalf("Unknown path element %d: '%c'", n, lr)
			}
		}

		// Check if we have finished all nodes or not. Break if we have.
		stillInPlay := 0
		for i, n := range currentNodes {
			if finishedNodes[i] == 0 {
				if n[2] == 'Z' {
					finishedNodes[i] = count
				} else {
					stillInPlay++
				}
			}
		}
		if stillInPlay == 0 {
			break
		}
	}

	part2 := leastCommonMultiple(finishedNodes)

	return part2
}

func main() {
	var inputFile = flag.String("input", "day08-input.txt", "Problem input file")
	flag.Parse()
	path, nodes := loadData(*inputFile)
	part1 := runDay08Part1(path, nodes)
	part2 := runDay08Part2(path, nodes)
	fmt.Printf("%d, %d\n", part1, part2)
}
