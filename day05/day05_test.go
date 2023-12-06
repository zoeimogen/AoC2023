package main

import "testing"

func TestDay05(t *testing.T) {
	part1, part2 := runDay05("day05-test.txt")

	if part1 != 35 {
		t.Errorf("Part 1 test returned %d; want 35", part1)
	}

	if part2 != 46 {
		t.Errorf("Part 2 test returned %d; want 46", part2)
	}
}
