package main

import (
    "testing"
)

func TestDay1(t *testing.T) {
    got := Day1_1("input-files/day1-test1.txt")
    if got != 514579 {
        t.Errorf("Day1_1() = %d; want 514579", got)
    }
    got = Day1_2("input-files/day1-test1.txt")
    if got != 241861950 {
        t.Errorf("Day1_2() = %d; want 241861950", got)
    }
}
