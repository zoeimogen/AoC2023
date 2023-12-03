package main

import "testing"

func TestDay03(t *testing.T) {
	part1, part2 := runDay03("day03-test.txt")

	if part1 != 4361 {
		t.Errorf("Part 1 test returned %d; want 4361", part1)
	}

	if part2 != 467835 {
		t.Errorf("Part 2 test returned %d; want 467835", part2)
	}
}
