package main

import "testing"

func TestDay09(t *testing.T) {
	part1, part2 := runDay09("day09-test.txt")

	if part1 != 114 {
		t.Errorf("Part 1 test returned %d; want 114", part1)
	}

	if part2 != 2 {
		t.Errorf("Part 2 test returned %d; want 2", part2)
	}
}
