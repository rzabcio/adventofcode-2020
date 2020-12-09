package main

import (
	"fmt"
)

func Day9_1(filename string) int {
	fmt.Printf("")
	codes := inputSlInt(filename)
	faultCode := findFaultCode(codes[1:], codes[0])
	return faultCode
}

func Day9_2(filename string) int {
	codes := inputSlInt(filename)
	faultCode := findFaultCode(codes[1:], codes[0])
	setOfAddends := findSetOfAddends(codes[1:], faultCode)
	return setOfAddends
}

func findFaultCode(codes []int, preLen int) int {
	for pos := preLen; pos < len(codes); pos++ {
		addPos1, _ := findAddends(codes[pos-preLen:pos], codes[pos])
		if addPos1 == -1 {
			return codes[pos]
		}
	}
	return 0
}

func findAddends(codes []int, result int) (int, int) {
	for pos1, addend1 := range codes {
		pos2 := indexOfInt(codes, result-addend1)
		if pos2 == -1 {
			continue
		}
		return pos1, pos2
	}
	return -1, -1
}

func findSetOfAddends(codes []int, result int) int {
	for pos1 := 0; pos1 < len(codes); pos1++ {
		sum := codes[pos1]
		min, max := sum, sum
		for pos2 := pos1 + 1; pos2 < len(codes); pos2++ {
			sum += codes[pos2]
			if sum == result {
				min, max = minMax(codes[pos1 : pos2+1])
				return min + max
			}
			if sum > result {
				break
			}
		}
	}
	return -1
}
