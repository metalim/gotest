package main

import (
	"fmt"
	"slices"
	"sort"
	"time"
)

var grid1 = [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
var grid2 = [][]int{{3, 2}, {1, 0}}
var grid3 [][]int

func init() {
	M := 1000
	N := 1000
	grid3 = make([][]int, M)
	for i := range grid3 {
		grid3[i] = make([]int, N)
	}
	for i := range grid3 {
		if i == 0 {
			continue
		}
		grid3[i][N-1] = -1
	}
}

func main() {
	timeRightIntervals([][]int{{3, 4}, {2, 3}, {1, 2}})
	timeRightIntervals([][]int{{1, 4}, {2, 5}, {7, 9}})
}

func timeBinarySearch(s string, t rune) {
	start := time.Now()
	i, found := sort.Find(len(s), func(i int) int {
		return int(t) - int(s[i])
	})
	fmt.Printf("%d, %t in %s\n", i, found, time.Since(start))
}

func countNegatives(grid [][]int) int {
	var count int
	for _, row := range grid {
		for _, num := range row {
			if num < 0 {
				count++
			}
		}
	}
	return count
}

func timeCountNegatives(grid [][]int) {
	start := time.Now()
	ans := countNegatives(grid)
	fmt.Printf("%d in %s\n", ans, time.Since(start))
}

func rightIntervals(intervals [][]int) []int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	fmt.Println(intervals)
	right := make([]int, len(intervals))
	return right
}

func timeRightIntervals(intervals [][]int) {
	start := time.Now()
	ans := rightIntervals(intervals)
	fmt.Printf("%v in %s\n", ans, time.Since(start))
}
