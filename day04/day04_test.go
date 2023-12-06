package main

import "testing"

func TestDay04(t *testing.T) {
	part1, part2 := runDay04("day04-test.txt")

	if part1 != 13 {
		t.Errorf("Part 1 test returned %d; want 13", part1)
	}

	if part2 != 30 {
		t.Errorf("Part 2 test returned %d; want 30", part2)
	}
}
