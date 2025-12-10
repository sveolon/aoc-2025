package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
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

func part1(lines []string) int {
	res := 0
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
