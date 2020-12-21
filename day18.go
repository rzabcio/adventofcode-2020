package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day18_1(filename string) int {
	fmt.Printf("")
	calc := NewElvCalc(filename)
	return calc.Calculate1()
}

func Day18_2(filename string) int {
	calc := NewElvCalc(filename)
	return calc.Calculate2()
}

type ElvCalc struct {
	eqs []ElvEqua
	res int
}

func NewElvCalc(filename string) ElvCalc {
	calc := new(ElvCalc)
	calc.eqs = make([]ElvEqua, 0)
	calc.res = 0
	for eqStr := range inputCh(filename) {
		eq := NewElvEqua(eqStr)
		calc.eqs = append(calc.eqs, eq)
	}
	return *calc
}

func (calc *ElvCalc) Calculate1() int {
	calc.res = 0
	for _, eq := range calc.eqs {
		eq.Calculate1()
		//fmt.Printf("   %s = %d\n", eq.orig, eq.res)
		calc.res += eq.res
	}
	return calc.res
}

func (calc *ElvCalc) Calculate2() int {
	calc.res = 0
	for _, eq := range calc.eqs {
		eq.Calculate2()
		//fmt.Printf("   %s = %d\n", eq.orig, eq.res)
		calc.res += eq.res
	}
	return calc.res
}

type ElvEqua struct {
	orig  string
	curr  string
	res   int
	r_par *regexp.Regexp
}

func NewElvEqua(eqStr string) ElvEqua {
	eq := new(ElvEqua)
	eq.orig = eqStr
	eq.curr = eqStr
	return *eq
}

func (eq *ElvEqua) Calculate1() int {
	eq.res = calcPart1(eq.curr)
	return eq.res
}

func (eq *ElvEqua) Calculate2() int {
	eq.res = calcPart2(eq.curr)
	return eq.res
}

// CALC TOOLS
var r_par = regexp.MustCompile(`\(([^()]*)\)`)
var r_add = regexp.MustCompile(`(\d+ \+ \d+)`)

func calcPart1(s string) int {
	for reduceParenthesis(&s, calcAsString) {
	}
	return calc(s)
}

func calcPart2(s string) int {
	for reduceParenthesis(&s, calcPart2AsString) {
	}
	for reduceAddition(&s) {
	}
	return calc(s)
}

type calcFunc func(s string) string

func reduceParenthesis(s *string, f calcFunc) bool {
	r := r_par.FindStringSubmatch(*s)
	//fmt.Printf("   r: %s, len: %d\n", s, len(r))
	if len(r) == 0 {
		return false
	}
	newS := r_par.ReplaceAllStringFunc(*s, f)
	//fmt.Printf("   %s => %s\n", s, newS)
	if newS == *s {
		return false
	}
	*s = newS
	return true
}

func reduceAddition(s *string) bool {
	r := r_add.FindStringSubmatch(*s)
	//fmt.Printf("   r: %s, len: %d\n", s, len(r))
	if len(r) == 0 {
		return false
	}
	newS := r_add.ReplaceAllStringFunc(*s, calcAsString)
	//fmt.Printf("   %s => %s\n", s, newS)
	if newS == *s {
		return false
	}
	*s = newS
	return true
}

func calcAsString(s string) string {
	return fmt.Sprintf("%d", calc(strings.Trim(s, "()")))
}

func calcPart2AsString(s string) string {
	return fmt.Sprintf("%d", calcPart2(strings.Trim(s, "()")))
}

func calc(s string) int {
	result := 0
	action := "+"
	for _, e := range strings.Split(s, " ") {
		if val, err := strconv.Atoi(e); err == nil {
			if action == "*" {
				result *= val
			} else {
				result += val
			}
		} else {
			action = e
		}
	}
	return result
}
