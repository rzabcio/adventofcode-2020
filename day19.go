package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day19_1(filename string) int {
	fmt.Printf("")
	mv := NewMsgValidator(filename)
	for mv.ParseRules() {
		//fmt.Println("", mv.Print())
	}
	return mv.CountRule0()
}

func Day19_2(filename string) int {
	calc := NewMsgValidator(filename)
	return calc.CountRule0()
}

type MsgValidator struct {
	rules []MsgRule
	msgs  []string
}

func NewMsgValidator(filename string) MsgValidator {
	mv := new(MsgValidator)
	ch := inputCh(filename)

	// parsing rules
	mv.rules = make([]MsgRule, 134)
	for line := range ch {
		if len(line) == 0 {
			break
		}
		rule := NewMsgRule(line)
		mv.rules[rule.id] = rule
	}

	// parsing messages
	mv.msgs = make([]string, 0)
	for msg := range ch {
		mv.msgs = append(mv.msgs, msg)
	}

	return *mv
}

func (mv *MsgValidator) ParseRules() bool {
	wasChange := false
	for i := 0; i < len(mv.rules); i++ {
		if len(mv.rules[i].orig) == 0 {
			continue
		}
		if mv.rules[i].Parsed {
			//fmt.Printf("[%d] already parsed: %s\n", i, mv.rules[i].Print())
			continue
		}
		if !mv.canBeParsed(mv.rules[i]) {
			//fmt.Printf("[%d] can't be parsed: %s\n", i, mv.rules[i].Print())
			continue
		}
		//fmt.Printf("[%d] will be parsed: %s\n", i, mv.rules[i].Print())

		mv.rules[i].s = mv.replaceIdsWithRules(mv.rules[i].s)
		mv.rules[i].Parse()
		wasChange = true
	}
	return wasChange
}

func (mv *MsgValidator) canBeParsed(rule MsgRule) bool {
	for _, id := range rule.SubRules() {
		if !mv.rules[id].Parsed {
			return false
		}
	}
	return true
}

func (mv *MsgValidator) replaceIdsWithRules(s string) string {
	return r_numbers.ReplaceAllStringFunc(s, mv.replaceIdWithRule)
}

func (mv *MsgValidator) replaceIdWithRule(s string) string {
	id, _ := strconv.Atoi(s)
	return fmt.Sprintf("%s", mv.rules[id].s)
}

func (mv *MsgValidator) CountRule0() int {
	count := 0
	for _, msg := range mv.msgs {
		if mv.rules[0].Matches(msg) {
			count++
		}
	}
	return count
}

func (mv *MsgValidator) Print() string {
	s := "rules:\n"
	for i, rule := range mv.rules {
		if len(rule.s) == 0 {
			continue
		}
		s += fmt.Sprintf("   [%d] - %s\n", i, rule.Print())
	}
	return s
}

type MsgRule struct {
	id     int
	orig   string
	s      string
	Parsed bool
	Regexp regexp.Regexp
}

func NewMsgRule(s string) MsgRule {
	rule := new(MsgRule)
	parsed := r_rule.FindStringSubmatch(s)
	rule.id, _ = strconv.Atoi(parsed[1])
	rule.Parsed = false
	rule.orig = parsed[2]
	rule.s = parsed[2]
	rule.Parse()
	return *rule
}

func (rule *MsgRule) Matches(s string) bool {
	return rule.Regexp.MatchString(s)
}

func (rule *MsgRule) Parse() bool {
	if r_numbers.MatchString(rule.s) {
		return false
	}
	if strings.Contains(rule.s, "| ") {
		rule.s = "(" + strings.ReplaceAll(rule.s, " | ", "|") + ")"
	}
	rule.s = strings.ReplaceAll(rule.s, "\"", "")
	rule.s = strings.ReplaceAll(rule.s, " ", "")
	rule.Regexp = *regexp.MustCompile("^" + rule.s + "$")
	rule.Parsed = true
	return true
}

func (rule *MsgRule) SubRules() []int {
	subRules := make([]int, 0)
	parsed := r_numbers.FindAllStringSubmatch(rule.orig, -1)
	for _, subParsed := range parsed {
		ruleId, _ := strconv.Atoi(subParsed[0])
		subRules = append(subRules, ruleId)
	}
	return subRules
}

func (rule *MsgRule) Print() string {
	s := ""
	s += rule.orig
	if rule.Parsed {
		s += " => " + rule.s
	}
	return s
}

// GLOBAL VARS
var r_numbers = regexp.MustCompile(`(\d+)`)
var r_rule = regexp.MustCompile(`^(\d*): (.*)$`)

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
