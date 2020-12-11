package main

import (
	"fmt"
	"strings"
)

func Day11_1(filename string) int {
	fmt.Printf("")
	gos := NewGameOfSeats(filename)
	gos.Run1()
	return gos.count(gos.currSeats, "#")
}

func Day11_2(filename string) int {
	return 0
}

/*
	This puzzle looks like game of life, so the name.
*/
type gameOfSeats struct {
	origSeats []string // .=no seat, L=seat, #=occupied
	currSeats []string
	nextSeats []string
}

func NewGameOfSeats(seatsFile string) gameOfSeats {
	gos := new(gameOfSeats)
	gos.origSeats = make([]string, 0)

	for line := range inputCh(seatsFile) {
		gos.origSeats = append(gos.origSeats, line)
	}
	gos.currSeats = make([]string, len(gos.origSeats))

	copy(gos.currSeats, gos.origSeats)
	gos.Reset()
	return *gos
}

func (this *gameOfSeats) Reset() {
	copy(this.currSeats, this.origSeats)
}

func (this *gameOfSeats) Run1() bool {
	for this.nextRound1() {
	}
	return true
}

func (this *gameOfSeats) nextRound1() bool {
	wasChange := false
	this.nextSeats = make([]string, 0)
	for y, seatLine := range this.currSeats {
		newSeatLine := ""
		for x, char := range seatLine {
			if "." == string(char) {
				newSeatLine += "."
				continue
			} else if "L" == string(char) && this.countAdjacents(x, y, "#") == 0 {
				newSeatLine += "#"
				wasChange = true
			} else if "#" == string(char) && this.countAdjacents(x, y, "#") >= 4 {
				newSeatLine += "L"
				wasChange = true
			} else {
				newSeatLine += string(char)
			}
		}
		this.nextSeats = append(this.nextSeats, newSeatLine)
	}
	copy(this.currSeats, this.nextSeats)
	return wasChange
}

func (this *gameOfSeats) countAdjacents(x, y int, what string) int {
	count := this.count(this.getAdjacents(x, y), what)
	for _, char := range what {
		if string(char) == string(this.currSeats[y][x]) {
			count--
		}
	}
	return count
}

func (this *gameOfSeats) count(seats []string, what string) int {
	count := 0
	for _, seatLine := range seats {
		for _, char := range what {
			count += strings.Count(seatLine, string(char))
		}
	}
	return count
}

func (this *gameOfSeats) getAdjacents(x, y int) []string {
	_, x_start := minMax([]int{0, x - 1})
	_, y_start := minMax([]int{0, y - 1})
	x_end, _ := minMax([]int{x + 1, len(this.origSeats[0]) - 1})
	y_end, _ := minMax([]int{y + 1, len(this.origSeats) - 1})
	result := make([]string, 0)
	for j := y_start; j <= y_end; j++ {
		subResult := this.currSeats[j][x_start : x_end+1]
		result = append(result, subResult)
	}
	return result
}

func printSeats(seats []string) string {
	s := ""
	for j := 0; j < len(seats); j++ {
		s += seats[j] + "\n"
	}
	return s
}
