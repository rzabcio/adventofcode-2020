package main

import (
	"sort"
	"strconv"
	"strings"
)


func Day5_1(filename string) int {
	seatIds := parsePlaneSeats(filename)
	sort.Ints(seatIds)
	return seatIds[len(seatIds)-1]
}

func Day5_2(filename string) int {
	seatIds := parsePlaneSeats(filename)
	sort.Ints(seatIds)
	for seatId := seatIds[0]; seatId<seatIds[len(seatIds)-1]+1; seatId++ {
		if seatId+seatIds[0] != seatIds[seatId] {
			return seatId+seatIds[0]
		}
	}
	return 0
}
func parsePlaneSeats(filename string) []int {
	seatIds := make([]int, 0)
	for seat := range inputCh(filename) {
		seatB := strings.ReplaceAll(strings.ReplaceAll(seat, "F", "0"), "B", "1")
		seatB = strings.ReplaceAll(strings.ReplaceAll(seatB, "L", "0"), "R", "1")
		seatId, _ := strconv.ParseInt(seatB, 2, 0)
		//fmt.Printf("    %s => %s => %d\n", seat, seatB, seatId)
		seatIds = append(seatIds, int(seatId))
	}
	return seatIds
}
