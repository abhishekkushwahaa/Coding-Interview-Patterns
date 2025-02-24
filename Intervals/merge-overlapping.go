package main

import (
	"fmt"
	"sort"
)

// Time: O(nlogn)
func sortInterval(intervals [][]int) {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sortInterval(intervals)

	mergedIntervals := make([][]int, 0, len(intervals))
	mergedIntervals = append(mergedIntervals, intervals[0])

	for _, intervals := range intervals[1:] {
		if top := mergedIntervals[len(mergedIntervals)-1]; intervals[0] > top[1] {
			mergedIntervals = append(mergedIntervals, intervals)
		} else if intervals[1] > top[1] {
			top[1] = intervals[1]
		}
	}

	return mergedIntervals
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	mergeIntervals := merge(intervals)
	fmt.Println("Merged Overlapping Intervals:", mergeIntervals)
}
