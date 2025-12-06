package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var ranges [][2]int
	emptySeen := false
	res := 0
	for _, s := range lines {
		s = strings.TrimSpace(s)
		if s == "" {
			emptySeen = true
			continue
		}

		if emptySeen == false {
			parts := strings.SplitN(s, "-", 2)
			from, _ := strconv.Atoi(parts[0])
			to, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, [2]int{from, to})
			continue
		} else {
			n, _ := strconv.Atoi(s)
			for _, r := range ranges {
				from := r[0]
				to := r[1]
				if from <= n && n <= to {
					res++
					break
				}
			}
		}
	}
	return res
}

func part2(lines []string) int {
	return 0
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
