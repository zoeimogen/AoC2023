package main

import "testing"

func TestDay01(t *testing.T) {
	part1 := runDay01Part1("day01-test.txt")

	if part1 != 142 {
		t.Errorf("Part 1 test returned %d; want 142", part1)
	}

	part2 := runDay01Part2("day01-test2.txt")

	if part2 != 281 {
		t.Errorf("Part 2 test returned %d; want 281", part2)
	}
}
