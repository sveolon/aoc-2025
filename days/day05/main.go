package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func debugAddOne(x int, where string) int {
    y := x + 1
    if (y < x) != (x < 0) {
        fmt.Printf("!!! OVERFLOW at %s: %d + 1 wrapped around\n", where, x)
    }
    return y
}

type Interval struct {
    Start int
    End   int
}

type IntervalSet struct {
    items []Interval // always sorted by Start, and non-overlapping
}

func (s *IntervalSet) find(pos int) int {
    // returns index of first interval with Start > pos
    return sort.Search(len(s.items), func(i int) bool {
        return s.items[i].Start > pos
    })
}

func (s *IntervalSet) Add(start, end int) {
    if start > end {
        start, end = end, start
    }

    i := s.find(start)

    if i > 0 && s.items[i-1].End+1 >= start {
        i--
    }

    newStart, newEnd := start, end

    j := i
    for j < len(s.items) && s.items[j].Start <= newEnd+1 {
        if s.items[j].Start < newStart {
            newStart = s.items[j].Start
        }
        if s.items[j].End > newEnd {
            newEnd = s.items[j].End
        }
        j++
    }

    // IMPORTANT: use a fresh slice so we don't overwrite data we still need
    merged := make([]Interval, 0, len(s.items)-(j-i)+1)
    merged = append(merged, s.items[:i]...)
    merged = append(merged, Interval{newStart, newEnd})
    merged = append(merged, s.items[j:]...)
    s.items = merged
}

func (s *IntervalSet) Contains(x int) bool {
    i := sort.Search(len(s.items), func(i int) bool {
        return s.items[i].Start > x
    })

    if i == 0 {
        return false
    }

    iv := s.items[i-1]
    return x >= iv.Start && x <= iv.End
}

func (s *IntervalSet) All() []Interval {
    return s.items
}

// 318561521520057 is too low
func part2(lines []string) int {
	var ranges IntervalSet
	for _, s := range lines {
		if s == "" {
			break
		}

		parts := strings.SplitN(s, "-", 2)
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		ranges.Add(from, to)
	}
	res := 0
	for _, r := range ranges.All() {
		res += (r.End - r.Start + 1)
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
