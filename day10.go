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
	adapters := inputSlInt(filename)
	sort.Ints(adapters)
	count := 1
	fmt.Println("+o ", adapters)
	count += remNextReduntantAdapter(adapters)
	//count = countAdapterConnections(adapters)
	return count
}

func findDiffs(adapters []int) []int {
	diffs := make([]int, 4)
	prev := 0
	for _, adapter := range adapters {
		diffs[adapter-prev]++
		prev = adapter
	}
	diffs[3]++
	return diffs
}

func remNextReduntantAdapter(adapters []int) int {
	count := 0
	prev := 0
	if len(adapters) < 2 {
		return 0
	}
	foundInThrees := make([]int, 0)
	for i := 0; i < len(adapters)-1; i++ {
		//fmt.Printf("=====> %d/%d\n", i, len(adapters))
		fmt.Printf("=====> %d/%d: %d\n", i, len(adapters), adapters)
		if i+2 < len(adapters) && adapters[i]-prev < 3 && adapters[i+1]-prev < 3 && adapters[i+2]-prev == 3 {
			adapters2 := adapters[i+2:]
			count += 3 + remNextReduntantAdapter(adapters2)
			foundInThrees = append(foundInThrees, adapters[i+1])
		} else if adapters[i]-prev < 3 && adapters[i+1]-prev <= 3 {
			adapters2 := adapters[i+1:]
			count += 1 + remNextReduntantAdapter(adapters2)
		}
		prev = adapters[i]
	}
	return count
}

func countAdapterConnections(adapters []int) int {
	count := 1
	prev := 0
	if len(adapters) < 2 {
		return 0
	}
	for i := 0; i < len(adapters)-1; {
		fmt.Printf("--- i=%d: %d\n", i, adapters)
		if i+2 < len(adapters) && adapters[i+2]-prev == 3 {
			count *= 4 - 1
			fmt.Printf("    2. can reduce %d and %d from prev: %s and %d\n", adapters[i], adapters[i+1], adapters[i+1], prev, adapters[i:])
			i += 2
		} else if i+1 < len(adapters) && adapters[i+1]-prev <= 3 {
			count *= 2
			fmt.Printf("    2. can reduce %d from prev: %s and %d\n", adapters[i], prev, adapters[i:])
			i += 1
		} else {
			i += 1
		}
		if i < len(adapters) {
			prev = adapters[i-1]
		}
	}
	return count
}
