package main

import (
	"fmt"
	"strings"
)

func Day11_1(filename string) int {
	fmt.Printf("")
	gos := NewGameOfSeats(filename)
	gos.Run(gos.getAdjacents, 4)
	return gos.count(gos.currSeats, "#")
}

func Day11_2(filename string) int {
	fmt.Printf("")
	gos := NewGameOfSeats(filename)
	gos.Run(gos.getVisible, 5)
	return gos.count(gos.currSeats, "#")
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

func (this *gameOfSeats) Run(adjacentFunc adjFunc, maxSeatsAdj int) bool {
	for this.nextRound(adjacentFunc, maxSeatsAdj) {
	}
	return true
}

func (this *gameOfSeats) nextRound(adjacentFunc adjFunc, maxSeatsAdj int) bool {
	wasChange := false
	this.nextSeats = make([]string, 0)
	for y, seatLine := range this.currSeats {
		newSeatLine := ""
		for x, char := range seatLine {
			if "." == string(char) {
				newSeatLine += "."
				continue
			} else if "L" == string(char) && this.countAdjacents(adjacentFunc, x, y, "#") == 0 {
				newSeatLine += "#"
				wasChange = true
			} else if "#" == string(char) && this.countAdjacents(adjacentFunc, x, y, "#") >= maxSeatsAdj {
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

func (this *gameOfSeats) countAdjacents(adjacentFunc adjFunc, x, y int, what string) int {
	count := this.count(adjacentFunc(x, y), what)
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

type adjFunc func(x, y int) []string

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

func (this *gameOfSeats) getVisible(xs, ys int) []string {
	result := make([]string, 0)
	for dy := -1; dy <= 1; dy++ {
		row := ""
		for dx := -1; dx <= 1; dx++ {
			// fmt.Printf("dx,dy=%d,%d\n", dx, dy)
			for x, y := xs+dx, ys+dy; y >= 0 && y < len(this.currSeats) && x >= 0 && x < len(this.currSeats[0]); x, y = x+dx, y+dy {
				seat := string(this.currSeats[y][x])
				// fmt.Printf("    looking in (%d,%d): %s\n", x, y, seat)
				if "L" == seat || "#" == seat {
					row += seat
					break
				}
			}
			if len(row) < dx+2 {
				row += "."
			}
		}
		result = append(result, row)
	}
	return result
}

//  0123456789
// 0      .
// 1     .
// 2.   .
// 3 . .
// 4  x
// 5 . .
// 6.   .
// 7     .
// 8      .
// 9       .

func printSeats(seats []string) string {
	s := ""
	for j := 0; j < len(seats); j++ {
		s += seats[j] + "\n"
	}
	return s
}
