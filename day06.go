package main

import (
	"fmt"
)


func Day6_1(filename string) int {
	fmt.Printf("")
	answers := readAnswers(filename)
	count := 0
	for _, gr := range answers {
		count += len(gr)-1
	}
	return count
}

func Day6_2(filename string) int {
	fmt.Printf("")
	answers := readAnswers(filename)
	count := 0
	for _, gr := range answers {
		for que, ansCount := range gr {
			if que == "persons" {
				continue
			}
			if ansCount == gr["persons"] {
				count++
			}
		}
	}
	return count
}

func readAnswers(filename string) []map[string]int {
	answers := make([]map[string]int, 0)
	gr := map[string]int{"persons": 0}
	for line := range inputCh(filename) {
		if line == "" {
			answers = append(answers, gr)
			gr = map[string]int{"persons": 0}
		} else {
			gr["persons"]++
		}
		for _, c := range line {
			if _, ok := gr[string(c)]; ok {
				gr[string(c)]++
			} else {
				gr[string(c)] = 1
			}
		}
	}
	answers = append(answers, gr)
	return answers
}
