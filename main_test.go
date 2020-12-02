package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	got, want := Day1_1("input-files/day1-test1.txt"), 514579
	if got != want {
		t.Errorf("Day1_1() = %d; want %s", got, want)
	}
	got, want = Day1_1("input-files/day1-test1.txt"), 241861950
	if got != want {
		t.Errorf("Day1_2() = %d; want %s", got, want)
	}
}

func TestDay2(t *testing.T) {
	got, want := Day2_1("input-files/day2-test1.txt"), 2
	if got != want {
		t.Error("Day2_1() =%d; want %s", got, want)
	}
}
