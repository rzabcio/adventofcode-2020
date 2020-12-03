package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func main() {
	m := map[string]func(string) int{
		"day1_1": Day1_1,
		"day1_2": Day1_2,
		"day2_1": Day2_1,
		"day2_2": Day2_2,
		"day3_1": Day3_1,
		"day3_2": Day3_2,
	}

	var day = &cobra.Command{
		Use:	"day [day_no] [test_no] [filename]",
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


// DAY 1 //////////////////////////////////////////////////////////////////////
func Day1_1(filename string) int {
	for no1 := range inputChInt(filename) {
		for no2 := range inputChInt(filename) {
			if no1+no2 == 2020 {
				return no1 * no2
			}
		}
	}
	return 0
}

func Day1_2(filename string) int {
	for no1 := range inputChInt(filename) {
		for no2 := range inputChInt(filename) {
			for no3 := range inputChInt(filename) {
				if no1+no2+no3 == 2020 {
					return no1 * no2 * no3
				}
			}
		}
	}
	return 0
}

// DAY 2 //////////////////////////////////////////////////////////////////////
func Day2_1(filename string) int {
	r := regexp.MustCompile(`(\d*)-(\d*) ([a-z]): ([a-z]*)`)
	passwdCount := 0
	for line := range inputCh(filename) {
		parsed := r.FindStringSubmatch(line)
		min, _ := strconv.Atoi(parsed[1])
		max, _ := strconv.Atoi(parsed[2])
		char, passwd := parsed[3], parsed[4]
		charCount := strings.Count(passwd, char)
		if min <= charCount && charCount <= max {
			passwdCount++
		}
	}
	return passwdCount
}

func Day2_2(filename string) int {
	r := regexp.MustCompile(`(\d*)-(\d*) ([a-z]): ([a-z]*)`)
	passwdCount := 0
	for line := range inputCh(filename) {
		parsed := r.FindStringSubmatch(line)
		pos1, _ := strconv.Atoi(parsed[1])
		pos2, _ := strconv.Atoi(parsed[2])
		char, passwd := parsed[3], parsed[4]
		pos1Char, pos2Char := string(passwd[pos1-1]), string(passwd[pos2-1])
		if (pos1Char==char || pos2Char==char) && pos1Char!=pos2Char {
			passwdCount++
		}
	}
	return passwdCount
}


// DAY 2 //////////////////////////////////////////////////////////////////////
func Day3_1(filename string) int {
	terrain := inputSl(filename)
	return countTrees(terrain, 3, 1)
}

func Day3_2(filename string) int {
	terrain := inputSl(filename)
	slopes := [5]int{11, 31, 51, 71, 12}
	treeCountMul := 1
	for _, slope := range slopes {
		dx, dy := int(slope/10), slope%10
		treeCountMul *= countTrees(terrain, dx, dy)
	}
	return treeCountMul
}

func countTrees(terrain []string, dx, dy int) int {
	treeCount := 0
	for x, y := 0, 0; y < len(terrain); x, y = (x+dx)%len(terrain[0]), y+dy {
		if '#' == terrain[y][x] {
			treeCount++
		}
	}
	return treeCount
}
