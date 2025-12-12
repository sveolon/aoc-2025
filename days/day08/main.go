package main

import (
	"bufio"
	"fmt"
	"os"
	//"math"
	"sort"
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

type Dist struct {
	//dist float64
	dist int
	p1 int // index in lines
	p2 int // index in lines
}

func parse3(s string) (int, int, int, error) {
    var a, b, c int
    _, err := fmt.Sscanf(s, "%d,%d,%d", &a, &b, &c)
    return a, b, c, err
}

func part1(lines []string) int {
	N := len(lines)
	dists := make([]Dist, 0, N*N/2)
	for i := 0; i < N; i++ {
		a1, b1, c1, _ := parse3(lines[i])
		for j := i + 1; j < N; j++ {
			a2, b2, c2, _ := parse3(lines[j])
			/*
			dist := math.Sqrt( 
				math.Pow(2, float64(a1-a2)) + 
				math.Pow(2, float64(b1-b2)) + 
				math.Pow(2, float64(c1-c2)))
				*/
			dist := (a1-a2)*(a1-a2) + (b1-b2)*(b1-b2) + (c1-c2)*(c1-c2)
			dists = append(dists, Dist{dist, i, j})
		}
	}
	sort.Slice(dists, func(i, j int) bool {
		return dists[i].dist < dists[j].dist 
	})
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
