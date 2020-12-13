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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func indexOfInt(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

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
