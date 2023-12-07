package main

import "testing"

func TestDay07(t *testing.T) {
	part1, part2 := runDay07("day07-test.txt")

	if part1 != 6440 {
		t.Errorf("Part 1 test returned %d; want 6440", part1)
	}

	if part2 != 5905 {
		t.Errorf("Part 2 test returned %d; want 5905", part2)
	}
}
