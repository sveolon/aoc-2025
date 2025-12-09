package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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
	ops := strings.Fields(lines[len(lines)-1])
	N := len(ops)
	curr := make([]int, N)
	for j := 0; j < N; j++ {
		if ops[j] == "+" {
			curr[j] = 0
		} else {
			curr[j] = 1
		}
	}
	for i := 0; i < len(lines)-1; i++ {
		nums := strings.Fields(lines[i])
		for j := 0; j < N; j++ {
			n, _ := strconv.Atoi(nums[j])
			if ops[j] == "+" {
				curr[j] += n
			} else {
				curr[j] *= n
			}
		}
	}

	res := 0
	for _, n := range curr {
		res += n
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
