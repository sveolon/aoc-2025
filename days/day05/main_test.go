package main

import "testing"

const (
	testPart1Expected = 3 // TODO: replace with known sample answer
	testPart2Expected = 14 // TODO: replace with known sample answer
	customExpected = 13
)

func TestSampleInput(t *testing.T) {
	lines := mustReadLines("test.txt")

	if got := part1(lines); got != testPart1Expected {
		t.Fatalf("part1: got %v, want %v", got, testPart1Expected)
	}
	if got := part2(lines); got != testPart2Expected {
		t.Fatalf("part2: got %v, want %v", got, testPart2Expected)
	}

	lines = mustReadLines("custom2.txt")
	if got := part2(lines); got != customExpected {
		t.Fatalf("custom2: got %v, want %v", got, customExpected)
	}
}
