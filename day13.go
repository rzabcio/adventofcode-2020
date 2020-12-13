package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day13_1(filename string) int {
	fmt.Printf("")
	timetable := NewShuttleTimetable(filename, 939)
	clShuttle, clTime := timetable.GetClosestShuttle()
	//fmt.Printf("shuttle %d leaves at %d (int %d minutes)\n", clShuttle, clTime, clTime-timetable.StartTime)
	return clShuttle * (clTime - timetable.StartTime)
}

func Day13_2(filename string) int {
	fmt.Printf("")
	timetable := NewShuttleTimetable(filename, 939)
	clShuttle, clTime := timetable.GetClosestShuttle()
	return clShuttle * (clTime - timetable.StartTime)
}

type ShuttleTimetable struct {
	Shuttles  []int
	StartTime int
	CurrTime  int
}

func NewShuttleTimetable(timetableFile string, startTime int) ShuttleTimetable {
	timetable := new(ShuttleTimetable)
	timetable.StartTime = startTime
	timetable.Reset()
	timetableFileData := inputSl(timetableFile)
	timetable.StartTime, _ = strconv.Atoi(timetableFileData[0])
	timetable.Shuttles = make([]int, 0)
	for _, shuttleS := range strings.Split(timetableFileData[1], ",") {
		if "x" == shuttleS {
			continue
		}
		shuttle, _ := strconv.Atoi(shuttleS)
		timetable.Shuttles = append(timetable.Shuttles, shuttle)
	}
	timetable.Reset()
	return *timetable
}

func (this *ShuttleTimetable) Reset() {
	this.CurrTime = this.StartTime
}

func (this *ShuttleTimetable) GetClosestShuttle() (int, int) {
	shuttle := 0
	_, maxShuttle := minMax(this.Shuttles)
	for this.CurrTime = this.StartTime; this.CurrTime <= this.StartTime+maxShuttle; this.CurrTime++ {
		for _, shuttle = range this.Shuttles {
			if this.CurrTime%shuttle == 0 {
				return shuttle, this.CurrTime
			}
		}
	}
	return 0, 0
}
