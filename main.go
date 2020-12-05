package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)


func main() {
	m := map[string]func(string) int{
		"day1_1": Day1_1, "day1_2": Day1_2,
		"day2_1": Day2_1, "day2_2": Day2_2,
		"day3_1": Day3_1, "day3_2": Day3_2,
		"day4_1": Day4_1, "day4_2": Day4_2,
		"day5_1": Day5_1, "day5_2": Day5_2,
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


// DAY 3 //////////////////////////////////////////////////////////////////////
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


// DAY 4 //////////////////////////////////////////////////////////////////////
func Day4_1(filename string) int {
	validCount := 0
	for _, p := range parsePassports(filename) {
		// fmt.Println("--- passport: ", p)
		if validatePassport(p, false) {
			validCount++
		}
	}
	return validCount
}

func Day4_2(filename string) int {
	validCount := 0
	for _, p := range parsePassports(filename) {
		if validatePassport(p, true) {
			validCount++
		}
	}
	return validCount
}

func parsePassports(filename string) []map[string]string {
	passports := make([]map[string]string, 0)
	p := make(map[string]string)
	for line := range inputCh(filename) {
		if line == "" {
			passports = append(passports, p)
			p = make(map[string]string)
			continue
		}
		parsePassportLine(p, line)
	}
	passports = append(passports, p)
	return passports
}
func parsePassportLine(p map[string]string, line string) {
	for _, entry := range strings.Split(line, " ") {
		split := strings.Split(entry, ":")
		key, val := split[0], split[1]
		p[key] = val
	}
}
func validatePassport(p map[string]string, validateFields bool) bool {
	//fmt.Println("--- passport: ", p)
	_, hasCid := p["cid"]
	if len(p) < 8 && (len(p)!=7 || hasCid) {
		//fmt.Println("    NOT OK: fields")
		return false
	}
	if !validateFields {
		return true
	}
	for key, val := range p {
		v_reg := v_regs[key]
		if v_reg != nil && !v_reg.MatchString(val) {
				//fmt.Println("    NOT OK:  v_reg: ", key)
				return false
		}
		v_func := v_funcs[key]
		if v_func != nil && !v_func(val) {
				//fmt.Println("    NOT OK: v_func: ", key)
				return false
		}
	}
	return true
}
var v_regs = map[string]*regexp.Regexp {
	"byr": regexp.MustCompile(`^\d{4}$`),
	"iyr": regexp.MustCompile(`^\d{4}$`),
	"eyr": regexp.MustCompile(`^\d{4}$`),
	"hgt": regexp.MustCompile(`^(\d{2,4})(cm|in)$`),
	"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
	"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
	"pid": regexp.MustCompile(`^\d{9}$`),
}
var v_funcs = map[string]func(string) bool{
	"byr": func(s string) bool {
					i, _ := strconv.Atoi(s)
					return 1920<=i && i<=2002
				},
	"iyr": func(s string) bool {
					i, _ := strconv.Atoi(s)
					return 2010<=i && i<=2020
				},
	"eyr": func(s string) bool {
					i, _ := strconv.Atoi(s)
					return 2020<=i && i<=2030
				},
	"hgt": func(s string) bool {
					heightUnit := v_regs["hgt"].FindStringSubmatch(s)
					height, _ := strconv.Atoi(heightUnit[1])
					if heightUnit[2] == "in" {
						return 59<=height && height<=76
					}
					return 150<=height && height<=193
				},
}


// DAY 5 //////////////////////////////////////////////////////////////////////
func Day5_1(filename string) int {
	maxSeatId := 0
	for _, seatId := range parsePlaneSeats(filename) {
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	return maxSeatId
}

func Day5_2(filename string) int {
	seatIds := parsePlaneSeats(filename)
	sort.Ints(seatIds)
	for seatId := seatIds[0]; seatId<seatIds[len(seatIds)-1]+1; seatId++ {
		if seatId+seatIds[0] != seatIds[seatId] {
			return seatId+seatIds[0]
		}
	}
	return 0
}
func parsePlaneSeats(filename string) []int {
	seatIds := make([]int, 0)
	for seat := range inputCh(filename) {
		seatB := strings.ReplaceAll(strings.ReplaceAll(seat, "F", "0"), "B", "1")
		seatB = strings.ReplaceAll(strings.ReplaceAll(seatB, "L", "0"), "R", "1")
		seatId, _ := strconv.ParseInt(seatB, 2, 0)
		//fmt.Printf("    %s => %s => %d\n", seat, seatB, seatId)
		seatIds = append(seatIds, int(seatId))
	}
	return seatIds
}
