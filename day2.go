package main

import (
	"regexp"
	"strconv"
	"strings"
)


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

