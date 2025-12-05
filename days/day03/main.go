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

func maxDigit(line string, from int, to int) (int, int) {
	f := line[from]
	p := from
	for i:=p+1; i <= to; i++ {
		if line[i] > f {
			f = line[i]
			p = i
		}
	}
	return int(f-'0'), p
}
func part1(lines []string) int {
	result := 0
	for _, line := range(lines) {
		f, p := maxDigit(line, 0, len(line) - 2)
		s, _ := maxDigit(line, p+1, len(line) - 1)
		result += f*10+s
		fmt.Println(line, result, f, s, p)
	}
	return result
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
