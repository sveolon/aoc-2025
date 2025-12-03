package main

import (
	"bufio"
	"fmt"
	"os"
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
	cur := 50
	res := 0
	for _, line := range lines {
		dir := line[0] // 'L' or 'R'
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if dir == 'L' {
			cur += 100
			cur -= num
		} else {
			cur += num
		}
		cur %= 100
		if cur == 0 {
			res++
		}

	}
	return res
}

func part2(lines []string) int {
	cur := 50
	res := 0
	for _, line := range lines {
		dir := line[0] // 'L' or 'R'
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		res += num / 100
		num %= 100
		if dir == 'L' {
			cur += 100
			cur -= num
			if cur <= 100 && cur + num != 100 {
				res++
			}
		} else {
			cur += num
			if cur >= 100 && cur - num != 100 {
				res++
			} 
		}

		cur %= 100

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
