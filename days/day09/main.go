package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputPath := "input.txt"
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	lines := mustReadLines(inputPath)

	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func parse2(s string) (int, int,  error) {
    var a, b  int
    _, err := fmt.Sscanf(s, "%d,%d", &a, &b)
    return a, b, err
}

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func part1(lines []string) int {
	N := len(lines)

	res := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			x1, y1, _ := parse2(lines[i])
			x2, y2, _ := parse2(lines[j])
			rect := (abs(x1-x2)+1) * (abs(y1-y2)+1)
			if rect < 0 {
				rect = -rect
			}
			res = max(res, rect)
		}
	}
	return res
}

func part2(lines []string) int {
	res := 0
	return res
}

func mustReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open %s: %v", path, err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("failed to read %s: %v", path, err))
	}
	return lines
}
