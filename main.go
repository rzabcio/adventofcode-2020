package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "github.com/spf13/cobra"
)

func main() {
  m := map[string]func(string) int {
        "day1_1": Day1_1,
        "day1_2": Day1_2,
  }

  var day = &cobra.Command{
    Use: "day [day_no] [test_no] [filename]",
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

func inputLines(filename string) (ch chan string) {
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

func inputLinesInt(filename string) (ch chan int) {
    ch = make(chan int)
    go func() {
        for str := range inputLines(filename) {
            i, _ := strconv.Atoi(str)
	    ch <- i
	}
        close(ch)
    }()
    return ch
}

func Day1_1(filename string) int {
    for no1 := range inputLinesInt(filename) {
        for no2 := range inputLinesInt(filename) {
            if no1+no2 == 2020 {
                return no1*no2
            }
	}
    }
    return 0
}

func Day1_2(filename string) int {
    for no1 := range inputLinesInt(filename) {
        for no2 := range inputLinesInt(filename) {
            for no3 := range inputLinesInt(filename) {
                if no1+no2+no3 == 2020 {
                     return no1*no2*no3
                }
            }
	}
    }
    return 0
}

func main2() {
    fmt.Println("day 1.1: ", Day1_1("day1_1.txt"))
    fmt.Println("day 1.2: ", Day1_2("day1_1.txt"))
}
