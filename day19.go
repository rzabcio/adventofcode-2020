package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day19_1(filename string) int {
	fmt.Printf("")
	mv := NewMessValidator(filename)
	test()
	return mv.CountRule0()
}

func Day19_2(filename string) int {
	calc := NewMessValidator(filename)
	return calc.CountRule0()
}

type MessValidator struct {
	rules []MessRule
	msgs  []string
}

func NewMessValidator(filename string) MessValidator {
	mv := new(MessValidator)
	ch := inputCh(filename)

	// parsing rules
	mv.rules = make([]MessRule, 133)
	r_rule := regexp.MustCompile(`^(\d*): (.*)$`)
	for line := range ch {
		if len(line) == 0 {
			break
		}
		parsed := r_rule.FindStringSubmatch(line)
		i, _ := strconv.Atoi(parsed[1])
		rule := NewMessRule(parsed[2])
		mv.rules[i] = rule
	}

	// parsing messages
	mv.msgs = make([]string, 0)
	for msg := range ch {
		mv.msgs = append(mv.msgs, msg)
	}

	return *mv
}

func (mv *MessValidator) CountRule0() int {
	return 0
}

type MessRule struct {
	orig   string
	s      string
	Parsed bool
	Regexp regexp.Regexp
}

func NewMessRule(s string) MessRule {
	rule := new(MessRule)
	rule.Parsed = false
	rule.orig = s
	rule.s = s
	rule.Parse()
	return *rule
}

func (rule *MessRule) Parse() bool {
	if strings.Contains(rule.s, "| ") {
		return false
	}
	rule.s = strings.Trim(rule.s, "\"")
	rule.Regexp = *regexp.MustCompile("^" + rule.s + "$")
	rule.Parsed = true
	return true
}

// test functions
func test() {
	s := "baab"
	subRules := []string{"(ab|ba)", "(ab|ba)"}
	rule := ""

	// AND
	for _, subRule := range subRules {
		rule += subRule
	}

	// OR
	//rule += "("
	//rule += strings.Join(subRules, "|")
	//rule += ")"

	s = "ac"
	rule = "((ab|ba)|(cd|dc))"

	fmt.Printf("final string rule: '%s'\n", rule)
	r := regexp.MustCompile("^" + rule + "$")
	match := r.MatchString(s)
	fmt.Printf("does '%s' match %s? %t\n", s, rule, match)
}
