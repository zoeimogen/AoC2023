package main

import "testing"

func TestDay06(t *testing.T) {
	part1, part2 := runDay06("day06-test.txt")

	if part1 != 288 {
		t.Errorf("Part 1 test returned %d; want 288", part1)
	}

	if part2 != 71503 {
		t.Errorf("Part 2 test returned %d; want 71503", part2)
	}
}
