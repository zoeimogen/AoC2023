package main

import "testing"

func TestDay02(t *testing.T) {
	part1, part2 := runDay02("day02-test.txt")

	if part1 != 8 {
		t.Errorf("Part 1 test returned %d; want 8", part1)
	}

	if part2 != 2286 {
		t.Errorf("Part 2 test returned %d; want 2286", part2)
	}
}
