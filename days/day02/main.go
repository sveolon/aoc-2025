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

func order(i int) int {
	result := 0
	for ; i != 0; i /= 10 {
		result++
	}
	return result
}
func isInvalid(i int) bool {
	ord := order(i)
	if ord % 2 != 0 {
		return false
	}
	pow := 1;
	for j := 0; j < ord / 2; j++ {
		pow *= 10
	}
	first := i / pow
	sec := i - first * pow
	return first == sec
}

func part1(lines []string) int {
	res := 0
	parts := strings.Split(lines[0], ",")

	for _, part := range parts {
		tokens := strings.Split(part, "-")
		from, _ := strconv.Atoi(tokens[0])
		to, _ := strconv.Atoi(tokens[1])
		for i := from; i <= to; i++ {
			if isInvalid(i) {
				res+=i
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
