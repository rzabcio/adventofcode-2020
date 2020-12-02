package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"reflect"
	"strconv"
)

func main() {
	m := map[string]func(string) int{
		"day1_1": Day1_1,
		"day1_2": Day1_2,
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
	return 0
}

func Day2_2(filename string) int {
	return 0
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

func ChToSl(ch interface{}) interface{} {
	chv := reflect.ValueOf(ch)
	slv := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(ch).Elem()), 0, 0)
	for {
		v, ok := chv.Recv()
		if !ok {
			return slv.Interface()
		}
	slv = reflect.Append(slv, v)
	}
}
