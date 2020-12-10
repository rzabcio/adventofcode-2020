package main

import (
	"fmt"
	"sort"
)

func Day10_1(filename string) int {
	fmt.Printf("")
	adapters := inputSlInt(filename)
	sort.Ints(adapters)
	diffs := findDiffs(adapters)
	return diffs[1] * diffs[3]
}

func Day10_2(filename string) int {
	result := 0
	return result
}

func findDiffs(adapters *[]int) []int {
	diffs := make([]int, 4)
	prev := 0
	for _, adapter := range adapters {
		diffs[adapter-prev]++
		prev = adapter
	}
	diffs[3]++
	return diffs
}
