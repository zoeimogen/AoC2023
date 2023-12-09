package main

import "testing"

func TestDay08(t *testing.T) {
	path, nodes := loadData("day08-test1.txt")
	part1 := runDay08Part1(path, nodes)

	if part1 != 2 {
		t.Errorf("Part 1 test 1 returned %d; want 2", part1)
	}

	path, nodes = loadData("day08-test2.txt")
	part1 = runDay08Part1(path, nodes)

	if part1 != 6 {
		t.Errorf("Part 1 test 2 returned %d; want 6", part1)
	}

	path, nodes = loadData("day08-test3.txt")
	part2 := runDay08Part2(path, nodes)

	if part2 != 6 {
		t.Errorf("Part 2 test returned %d; want 6", part2)
	}
}
