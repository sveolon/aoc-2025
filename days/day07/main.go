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
	L := len(lines)
	C := len(lines[0])
	cur := make([]int, C)
	for c := 0; c < C; c++ {
		if lines[0][c] == 'S' {
			cur[c] = 1
		}
	}
	res := 0
	for l := 1; l < L; l++ {
		newCur := make([]int, C)
		for c := 0; c < C; c++ {
			if cur[c] == 1 {
				switch lines[l][c] {
				case '.':
						newCur[c] = 1
					case '^':
						res++
						if c > 0 {
							newCur[c-1] = 1
						}
						if c < C-1 {
							newCur[c+1] = 1
						}
				} 
			}
		}
		cur = newCur
	}
	return res
}

func part2(lines []string) int {
	L := len(lines)
	C := len(lines[0])
	cur := make([]int, C)
	for c := 0; c < C; c++ {
		if lines[0][c] == 'S' {
			cur[c] = 1
		}
	}
	for l := 1; l < L; l++ {
		newCur := make([]int, C)
		for c := 0; c < C; c++ {
			if cur[c] > 0 {
				switch lines[l][c] {
				case '.':
						newCur[c] += cur[c]
					case '^':
						if c > 0 {
							newCur[c-1] += cur[c]
						}
						if c < C-1 {
							newCur[c+1] += cur[c]
						}
				} 
			}
		}
		cur = newCur
	}
	res := 0
	for c := 0; c < C; c++ {
		res += cur[c]
	}
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
