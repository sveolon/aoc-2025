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

	fmt.Printf("Part 1: %v\n", part1(lines, 1000))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

type DisjointSetElement struct {
	Parent int // parent index; root has Parent == itself
	Size   int // size of the set; valid only for roots
}

type DisjointSet struct {
	Set []DisjointSetElement
}

// NewDisjointSet creates a DSU with elements 0..n-1,
// each in its own singleton set of size 1.
func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		Set: make([]DisjointSetElement, n),
	}
	for i := range ds.Set {
		ds.Set[i] = DisjointSetElement{
			Parent: i,
			Size:   1,
		}
	}
	return ds
}

// Find returns the root of the set containing x,
// with path compression to flatten the tree.
func (ds *DisjointSet) Find(x int) int {
	if ds.Set[x].Parent != x {
		ds.Set[x].Parent = ds.Find(ds.Set[x].Parent)
	}
	return ds.Set[x].Parent
}

// Union merges the sets containing x and y using union by size:
// the smaller set is attached under the larger one.
// It returns the root of the merged set.
func (ds *DisjointSet) Union(x, y int) int {
	rx := ds.Find(x)
	ry := ds.Find(y)

	if rx == ry {
		return rx // already in the same set
	}

	// Ensure rx is the root of the larger set
	if ds.Set[rx].Size < ds.Set[ry].Size {
		rx, ry = ry, rx
	}

	// Attach root ry under root rx
	ds.Set[ry].Parent = rx
	ds.Set[rx].Size += ds.Set[ry].Size
	//ds.Set[ry].Size = 0

	return rx
}

// SameSet reports whether x and y are in the same set.
func (ds *DisjointSet) SameSet(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}

// SizeOf returns the size of the set containing x.
func (ds *DisjointSet) SizeOf(x int) int {
	root := ds.Find(x)
	return ds.Set[root].Size
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

// ThreeLargestComponents returns up to three largest components by size.
// It returns arrays of roots and sizes; unused slots (if fewer than 3 components)
// will have root == -1 and size == 0.
func (ds *DisjointSet) ThreeLargestComponents() (roots [3]int, sizes [3]int) {
    // Optional: compress paths so every node points directly to its root.
    // This doesn't change sizes, just makes future operations cheaper.
    for i := range ds.Set {
        ds.Find(i)
    }

    // Initialize roots to -1 to mark "empty"
    roots = [3]int{-1, -1, -1}
    sizes = [3]int{0, 0, 0}

    for i, el := range ds.Set {
        // Only consider roots
        if el.Parent != i {
            continue
        }
        s := el.Size

        // Insert into top-3 (simple fixed-size "leaderboard")
        switch {
        case s > sizes[0]:
            // shift down
            sizes[2], roots[2] = sizes[1], roots[1]
            sizes[1], roots[1] = sizes[0], roots[0]
            sizes[0], roots[0] = s, i
        case s > sizes[1]:
            sizes[2], roots[2] = sizes[1], roots[1]
            sizes[1], roots[1] = s, i
        case s > sizes[2]:
            sizes[2], roots[2] = s, i
        }
    }

    return roots, sizes
}

func part1(lines []string, connections int) int {
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

	ds := NewDisjointSet(N)
	for i, c := 0, 0; i < connections; i++ {
		conn := dists[i]
		if !ds.SameSet(conn.p1, conn.p2) {
			c++

			ds.Union(conn.p1, conn.p2)
		}
	}

	res := 1
	roots, sizes := ds.ThreeLargestComponents()
	for i := 0; i < 3; i++ {
		if roots[i] == -1 {
			continue
		}
		res *= sizes[i]
	}

	return res
}

func part2(lines []string) int {
	N := len(lines)
	dists := make([]Dist, 0, N*N/2)
	for i := 0; i < N; i++ {
		a1, b1, c1, _ := parse3(lines[i])
		for j := i + 1; j < N; j++ {
			a2, b2, c2, _ := parse3(lines[j])
			dist := (a1-a2)*(a1-a2) + (b1-b2)*(b1-b2) + (c1-c2)*(c1-c2)
			dists = append(dists, Dist{dist, i, j})
		}
	}
	sort.Slice(dists, func(i, j int) bool {
		return dists[i].dist < dists[j].dist 
	})

	ds := NewDisjointSet(N)
	n := N-1
	for i := 0; ; i++ {
		conn := dists[i]
		if !ds.SameSet(conn.p1, conn.p2) {
			if n == 1 {
				x1, _, _, _ := parse3(lines[conn.p1])
				x2, _, _, _ := parse3(lines[conn.p2])
				return x1 * x2
			}
			n--
			ds.Union(conn.p1, conn.p2)
		}
	}
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
