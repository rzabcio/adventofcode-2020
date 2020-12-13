package main

import (
	"fmt"
	"sort"
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
	clTime := timetable.GetClosestCombo2()
	return clTime
}

type ShuttleTimetable struct {
	Shuttles  []Shuttle
	StartTime int
	CurrTime  int
}

type Shuttle struct {
	no  int
	ind int
	rem int
}

func NewShuttleTimetable(timetableFile string, startTime int) ShuttleTimetable {
	this := new(ShuttleTimetable)
	this.StartTime = startTime
	this.Reset()
	timetableFileData := inputSl(timetableFile)
	this.StartTime, _ = strconv.Atoi(timetableFileData[0])
	this.Shuttles = make([]Shuttle, 0)
	for i, ss := range strings.Split(timetableFileData[1], ",") {
		if "x" != ss {
			shuttle := new(Shuttle)
			shuttle.ind, shuttle.rem = i, i
			shuttle.no, _ = strconv.Atoi(ss)
			if shuttle.ind > 0 {
				shuttle.rem = shuttle.no - (shuttle.ind % shuttle.no)
			}
			this.Shuttles = append(this.Shuttles, *shuttle)
		}
	}
	sortShuttles(&this.Shuttles)
	this.Reset()
	return *this
}

func (this *ShuttleTimetable) Reset() {
	this.CurrTime = this.StartTime
}

func (this *ShuttleTimetable) GetClosestShuttle() (int, int) {
	maxShuttle := this.Shuttles[0]
	for this.CurrTime = this.StartTime; this.CurrTime <= this.StartTime+maxShuttle.no; this.CurrTime++ {
		for _, shuttle := range this.Shuttles {
			if this.CurrTime%shuttle.no == 0 {
				return shuttle.no, this.CurrTime
			}
		}
	}
	return 0, 0
}

func (this *ShuttleTimetable) GetClosestCombo2() int {
	max := this.Shuttles[0].no
	maxIndex := this.Shuttles[0].ind
	for this.CurrTime = this.StartTime; ; this.CurrTime++ {
		if (this.CurrTime+maxIndex)%max == 0 {
			break
		}
	}
	deltaTime := 1
	for i := 0; i < len(this.Shuttles)-1; i++ {
		shuttle := this.Shuttles[i]
		deltaTime *= shuttle.no
		for this.Shuttles[i+1].rem != this.CurrTime%this.Shuttles[i+1].no {
			this.CurrTime += deltaTime
		}
	}
	return this.CurrTime
}

func sortShuttles(shuttles *[]Shuttle) {
	newShuttles := make([]Shuttle, 0)
	shuttleInts := make([]int, 0)
	for _, shuttle := range *shuttles {
		shuttleInts = append(shuttleInts, shuttle.no)
	}
	sort.Ints(shuttleInts)
	for i := len(shuttleInts) - 1; i >= 0; i-- {
		for _, shuttle := range *shuttles {
			if shuttleInts[i] == shuttle.no {
				newShuttles = append(newShuttles, shuttle)
			}
		}
	}
	copy(*shuttles, newShuttles)
}
