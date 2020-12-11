package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	got, want := Day1_1("input-files/day1-test1.txt"), 514579
	if got != want {
		t.Errorf("Day1_1() = %d; want %d", got, want)
	}
	got, want = Day1_2("input-files/day1-test1.txt"), 241861950
	if got != want {
		t.Errorf("Day1_2() = %d; want %d", got, want)
	}
}

func TestDay2(t *testing.T) {
	got, want := Day2_1("input-files/day2-test1.txt"), 2
	if got != want {
		t.Errorf("Day2_1() = %d; want %d", got, want)
	}
	got, want = Day2_2("input-files/day2-test1.txt"), 1
	if got != want {
		t.Errorf("Day2_2() = %d; want %d", got, want)
	}
}

func TestDay3(t *testing.T) {
	got, want := Day3_1("input-files/day3-test1.txt"), 7
	if got != want {
		t.Errorf("Day3_1() = %d; want %d", got, want)
	}
	got, want = Day3_2("input-files/day3-test1.txt"), 336
	if got != want {
		t.Errorf("Day3_2() = %d; want %d", got, want)
	}
}

func TestDay4(t *testing.T) {
	got, want := Day4_1("input-files/day4-test1.txt"), 2
	if got != want {
		t.Errorf("Day4_1(test1) = %d; want %d", got, want)
	}
	got, want = Day4_2("input-files/day4-test1.txt"), 2
	if got != want {
		t.Errorf("Day4_2(test1) = %d; want %d", got, want)
	}
	got, want = Day4_2("input-files/day4-test-valids.txt"), 4
	if got != want {
		t.Errorf("Day4_2(valids) = %d; want %d", got, want)
	}
	got, want = Day4_2("input-files/day4-test-invalids.txt"), 0
	if got != want {
		t.Errorf("Day4_2(invalids) = %d; want %d", got, want)
	}
}

func TestDay5(t *testing.T) {
	got, want := Day5_1("input-files/day5-seats-test1.txt"), 820
	if got != want {
		t.Errorf("Day5_1(test1) = %d; want %d", got, want)
	}
}

func TestDay6(t *testing.T) {
	got, want := Day6_1("input-files/day6-answers-test1.txt"), 11
	if got != want {
		t.Errorf("Day6_1(test1) = %d; want %d", got, want)
	}
	got, want = Day6_2("input-files/day6-answers-test1.txt"), 6
	if got != want {
		t.Errorf("Day6_2(test1) = %d; want %d", got, want)
	}
}

func TestDay7(t *testing.T) {
	got, want := Day7_1("input-files/day7-bagrules-test1.txt"), 4
	if got != want {
		t.Errorf("Day7_1(test1) = %d; want %d", got, want)
	}
	got, want = Day7_2("input-files/day7-bagrules-test1.txt"), 32
	if got != want {
		t.Errorf("Day7_2(test1) = %d; want %d", got, want)
	}
	got, want = Day7_2("input-files/day7-bagrules-test2.txt"), 126
	if got != want {
		t.Errorf("Day7_2(test1) = %d; want %d", got, want)
	}
}

func TestDay8(t *testing.T) {
	got, want := Day8_1("input-files/day8-program-test1.txt"), 5
	if got != want {
		t.Errorf("Day8_1(test1) = %d; want %d", got, want)
	}
	got, want = Day8_2("input-files/day8-program-test1.txt"), 8
	if got != want {
		t.Errorf("Day8_2(test1) = %d; want %d", got, want)
	}
}

func TestDay9(t *testing.T) {
	got, want := Day9_1("input-files/day9-codes-test1.txt"), 127
	if got != want {
		t.Errorf("Day9_1(test1) = %d; want %d", got, want)
	}
	got, want = Day9_2("input-files/day9-codes-test1.txt"), 62
	if got != want {
		t.Errorf("Day9_2(test1) = %d; want %d", got, want)
	}
}

func TestDay10(t *testing.T) {
	got, want := Day10_1("input-files/day10-adapters-test1.txt"), 35
	if got != want {
		t.Errorf("Day10_1(test1) = %d; want %d", got, want)
	}
	got, want = Day10_1("input-files/day10-adapters-test2.txt"), 220
	if got != want {
		t.Errorf("Day10_1(test2) = %d; want %d", got, want)
	}

	got, want = Day10_2("input-files/day10-adapters-test1.txt"), 8
	if got != want {
		t.Errorf("Day10_2(test1) = %d; want %d", got, want)
	}
	got, want = Day10_2("input-files/day10-adapters-test2.txt"), 19208
	if got != want {
		t.Errorf("Day10_2(test2) = %d; want %d", got, want)
	}
}
