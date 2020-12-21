package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	got, want := Day1_1("input-files/day01-test1.txt"), 514579
	if got != want {
		t.Errorf("Day1_1() = %d; want %d", got, want)
	}
	got, want = Day1_2("input-files/day01-test1.txt"), 241861950
	if got != want {
		t.Errorf("Day1_2() = %d; want %d", got, want)
	}
}

func TestDay2(t *testing.T) {
	got, want := Day2_1("input-files/day02-test1.txt"), 2
	if got != want {
		t.Errorf("Day2_1() = %d; want %d", got, want)
	}
	got, want = Day2_2("input-files/day02-test1.txt"), 1
	if got != want {
		t.Errorf("Day2_2() = %d; want %d", got, want)
	}
}

func TestDay3(t *testing.T) {
	got, want := Day3_1("input-files/day03-test1.txt"), 7
	if got != want {
		t.Errorf("Day3_1() = %d; want %d", got, want)
	}
	got, want = Day3_2("input-files/day03-test1.txt"), 336
	if got != want {
		t.Errorf("Day3_2() = %d; want %d", got, want)
	}
}

func TestDay4(t *testing.T) {
	got, want := Day4_1("input-files/day04-test1.txt"), 2
	if got != want {
		t.Errorf("Day4_1(test1) = %d; want %d", got, want)
	}
	got, want = Day4_2("input-files/day04-test1.txt"), 2
	if got != want {
		t.Errorf("Day4_2(test1) = %d; want %d", got, want)
	}
	got, want = Day4_2("input-files/day04-test-valids.txt"), 4
	if got != want {
		t.Errorf("Day4_2(valids) = %d; want %d", got, want)
	}
	got, want = Day4_2("input-files/day04-test-invalids.txt"), 0
	if got != want {
		t.Errorf("Day4_2(invalids) = %d; want %d", got, want)
	}
}

func TestDay5(t *testing.T) {
	got, want := Day5_1("input-files/day05-seats-test1.txt"), 820
	if got != want {
		t.Errorf("Day5_1(test1) = %d; want %d", got, want)
	}
}

func TestDay6(t *testing.T) {
	got, want := Day6_1("input-files/day06-answers-test1.txt"), 11
	if got != want {
		t.Errorf("Day6_1(test1) = %d; want %d", got, want)
	}
	got, want = Day6_2("input-files/day06-answers-test1.txt"), 6
	if got != want {
		t.Errorf("Day6_2(test1) = %d; want %d", got, want)
	}
}

func TestDay7(t *testing.T) {
	got, want := Day7_1("input-files/day07-bagrules-test1.txt"), 4
	if got != want {
		t.Errorf("Day7_1(test1) = %d; want %d", got, want)
	}
	got, want = Day7_2("input-files/day07-bagrules-test1.txt"), 32
	if got != want {
		t.Errorf("Day7_2(test1) = %d; want %d", got, want)
	}
	got, want = Day7_2("input-files/day07-bagrules-test2.txt"), 126
	if got != want {
		t.Errorf("Day7_2(test1) = %d; want %d", got, want)
	}
}

func TestDay8(t *testing.T) {
	got, want := Day8_1("input-files/day08-program-test1.txt"), 5
	if got != want {
		t.Errorf("Day8_1(test1) = %d; want %d", got, want)
	}
	got, want = Day8_2("input-files/day08-program-test1.txt"), 8
	if got != want {
		t.Errorf("Day8_2(test1) = %d; want %d", got, want)
	}
}

func TestDay9(t *testing.T) {
	got, want := Day9_1("input-files/day09-codes-test1.txt"), 127
	if got != want {
		t.Errorf("Day9_1(test1) = %d; want %d", got, want)
	}
	got, want = Day9_2("input-files/day09-codes-test1.txt"), 62
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
}

func TestDay11(t *testing.T) {
	got, want := Day11_1("input-files/day11-seats-test1.txt"), 37
	if got != want {
		t.Errorf("Day11_1(test1) = %d; want %d", got, want)
	}

	got, want = Day11_2("input-files/day11-seats-test1.txt"), 26
	if got != want {
		t.Errorf("Day11_2(test1) = %d; want %d", got, want)
	}
}

func TestDay12(t *testing.T) {
	got, want := Day12_1("input-files/day12-navigations-test.txt"), 25
	if got != want {
		t.Errorf("Day12_1(test1) = %d; want %d", got, want)
	}

	got, want = Day12_2("input-files/day12-navigations-test.txt"), 286
	if got != want {
		t.Errorf("Day12_2(test1) = %d; want %d", got, want)
	}
}

func TestDay13(t *testing.T) {
	got, want := Day13_1("input-files/day13-shuttles-test1.txt"), 295
	if got != want {
		t.Errorf("Day13_1(test1) = %d; want %d", got, want)
	}

	got, want = Day13_2("input-files/day13-shuttles-test1.txt"), 1068781
	if got != want {
		t.Errorf("Day13_2(test1) = %d; want %d", got, want)
	}
	got, want = Day13_2("input-files/day13-shuttles-test2.txt"), 754018
	if got != want {
		t.Errorf("Day13_2(test1) = %d; want %d", got, want)
	}

}

func TestDay14(t *testing.T) {
	got, want := Day14_1("input-files/day14-docking-test1.txt"), 165
	if got != want {
		t.Errorf("Day14_1(test1) = %d; want %d", got, want)
	}

	got, want = Day14_2("input-files/day14-docking-test2.txt"), 208
	if got != want {
		t.Errorf("Day14_2(test1) = %d; want %d", got, want)
	}
}

func TestDay15(t *testing.T) {
	got, want := Day15_1("input-files/day15-startingNumbers-test1.txt"), 436
	if got != want {
		t.Errorf("Day15_1(test1) = %d; want %d", got, want)
	}

	got, want = Day15_1("input-files/day15-startingNumbers-test2.txt"), 1836
	if got != want {
		t.Errorf("Day15_1(test1) = %d; want %d", got, want)
	}
}

func TestDay16(t *testing.T) {
	got, want := Day16_1("input-files/day16-tickets-test1.txt"), 71
	if got != want {
		t.Errorf("Day16_1(test1) = %d; want %d", got, want)
	}

	got, want = Day16_2("input-files/day16-tickets.txt"), 453459307723
	if got != want {
		t.Errorf("Day16_2(test1) = %d; want %d", got, want)
	}
}

//func TestDay17(t *testing.T) {
//	got, want := Day17_1("input-files/day17-cubespace-test.txt"), 112
//	if got != want {
//		t.Errorf("Day17_1(test1) = %d; want %d", got, want)
//	}
//
//	got, want = Day17_2("input-files/day17-cubespace-test.txt"), 453459307723
//	if got != want {
//		t.Errorf("Day17_2(test1) = %d; want %d", got, want)
//	}
//}

func TestDay18(t *testing.T) {
	got, want := Day18_1("input-files/day18-homework-test1.txt"), 26457
	if got != want {
		t.Errorf("Day18_1(test1) = %d; want %d", got, want)
	}

	got, want = Day18_2("input-files/day18-homework-test1.txt"), 694173
	if got != want {
		t.Errorf("Day18_2(test1) = %d; want %d", got, want)
	}
}
