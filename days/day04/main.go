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
	L := len(lines)
	C := len(lines[0])
	for l := 0; l < L; l++ {
		for c := 0; c < C; c++ {
			if lines[l][c] != '@' {
				continue
			}
			count := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if l + i < 0 || l + i >= L || c + j < 0 || c + j >= C || (i == 0 && j == 0) {
						continue
					}
					if lines[l+i][c+j] == '@' {
						count++
					}
				}
			}
			if count < 4 {
				res++
			}
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
