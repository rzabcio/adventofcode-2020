package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func main() {
	m := map[string]func(string) int{
		"day1_1": Day1_1, "day1_2": Day1_2,
		"day2_1": Day2_1, "day2_2": Day2_2,
		"day3_1": Day3_1, "day3_2": Day3_2,
		"day4_1": Day4_1, "day4_2": Day4_2,
		"day5_1": Day5_1, "day5_2": Day5_2,
		"day6_1": Day6_1, "day6_2": Day6_2,
		"day7_1": Day7_1, "day7_2": Day7_2,
		"day8_1": Day8_1, "day8_2": Day8_2,
		"day9_1": Day9_1, "day9_2": Day9_2,
		"day10_1": Day10_1, "day10_2": Day10_2,
		"day11_1": Day11_1, "day11_2": Day11_2,
		"day12_1": Day12_1, "day12_2": Day12_2,
		"day13_1": Day13_1, "day13_2": Day13_2,
		"day14_1": Day14_1, "day14_2": Day14_2,
		"day15_1": Day15_1, "day15_2": Day15_2,
		"day16_1": Day16_1, "day16_2": Day16_2,
		"day17_1": Day17_1, "day17_2": Day17_2,
		"day18_1": Day18_1, "day18_2": Day18_2,
		"day19_1": Day19_1, "day19_2": Day19_2,
		"day20_1": Day20_1, "day20_2": Day20_2,
		"day21_1": Day21_1, "day21_2": Day21_2,
	}

	var day = &cobra.Command{
		Use:  "day [day_no] [test_no] [filename]",
		Args: cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			f := m["day"+args[0]+"_"+args[1]]
			fmt.Println(f(args[2]))
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(day)
	rootCmd.Execute()
}

// TOOLS //////////////////////////////////////////////////////////////////////
func inputSl(filename string) []string {
	sl := make([]string, 0)
	for s := range inputCh(filename) {
		sl = append(sl, s)
	}
	return sl
}

func inputSlInt(filename string) []int {
	sl := make([]int, 0)
	for s := range inputChInt(filename) {
		sl = append(sl, s)
	}
	return sl
}

func inputCh(filename string) (ch chan string) {
	ch = make(chan string)
	go func() {
		//file, err := os.Open("input-files/"+filename)
		file, err := os.Open(filename)
		if err != nil {
			close(ch)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch
}

func inputChInt(filename string) (ch chan int) {
	ch = make(chan int)
	go func() {
		for str := range inputCh(filename) {
			i, _ := strconv.Atoi(str)
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// TOOLS - STRING
func ReverseStr(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

// TOOLS - ARRAYS
func remove(s []string, e string) []string {
	i := indexOf(s, e)
	if i < 0 {
		return s
	}
	res := make([]string, 0)
	if i == 0 {
		res = s[i+1:]
	} else if i == len(s)-1 {
		res = s[:i]
	} else {
		res = append(s[:i], s[i+1:]...)
	}
	return res
}

func contains(s []string, e string) bool {
	return indexOf(s, e) >= 0
}

func indexOf(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func containsInt(s []int, e int) bool {
	return indexOfInt(s, e) >= 0
}

func indexOfInt(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func ReverseStrArr(ss []string) []string {
	for i := 0; i < len(ss)/2; i++ {
		j := len(ss) - i - 1
		ss[i], ss[j] = ss[j], ss[i]
	}
	return ss
}

func intersection(a []string, b []string) (inter []string) {
	// interacting on the smallest list first can potentailly be faster...but not by much, worse case is the same
	low, high := a, b
	if len(a) > len(b) {
		low = b
		high = a
	}

	done := false
	for i, l := range low {
		for j, h := range high {
			// get future index values
			f1 := i + 1
			f2 := j + 1
			if l == h {
				inter = append(inter, h)
				if f1 < len(low) && f2 < len(high) {
					// if the future values aren't the same then that's the end of the intersection
					if low[f1] != high[f2] {
						done = true
					}
				}
				// we don't want to interate on the entire list everytime, so remove the parts we already looped on will make it faster each pass
				high = high[:j+copy(high[j:], high[j+1:])]
				break
			}
		}
		// nothing in the future so we are done
		if done {
			break
		}
	}
	return
}

// TOOLS - NUMERICAL
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
