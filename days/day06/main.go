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
	L := len(lines)
	C := len(lines[0])
	tmp := 0
	res := 0
	isMult := false
	for c := 0; c < C; c++ {
		if lines[L-1][c] != ' ' {
			res += tmp;
			isMult = (lines[L-1][c] == '*')
			if isMult {
				tmp = 1
			} else {
				tmp = 0
			}
		}
		num := 0
		numC := 0
		for l := 0; l < L-1; l++ {
			if lines[l][c] != ' ' {
				numC++
				num *= 10
				n := lines[l][c] - '0'
				num += int(n)
			}
		}
		if numC > 0 {
			if isMult{
				tmp *= num
			} else {
				tmp += num
			}
		}
	}
	res += tmp
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
