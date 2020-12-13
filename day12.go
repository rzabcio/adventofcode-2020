package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day12_1(filename string) int {
	fmt.Printf("")
	vessel := NewVessel()
	vessel.SetSail(filename, vessel.sail)
	return vessel.DistFromStart()
}

func Day12_2(filename string) int {
	fmt.Printf("")
	vessel := NewVessel()
	vessel.SetSail(filename, vessel.sailToWaypoint)
	return vessel.DistFromStart()
}

type Vessel struct {
	Pos_x      int
	Pos_y      int
	Waypoint_x int
	Waypoint_y int
	facings    []string
	currFacing int

	reg_course *regexp.Regexp
}

func NewVessel() Vessel {
	vessel := new(Vessel)
	vessel.Reset()
	vessel.facings = []string{"N", "E", "S", "W"}
	vessel.reg_course = regexp.MustCompile(`^([A-Z])(\d*)$`)
	return *vessel
}

func (this *Vessel) Reset() {
	this.Pos_x, this.Pos_y = 0, 0
	this.Waypoint_x, this.Waypoint_y = 10, 1
	this.currFacing = 1
}

func (this *Vessel) DistFromStart() int {
	return AbsInt(this.Pos_x) + AbsInt(this.Pos_y)
}

func (this *Vessel) SetSail(navFile string, sailFunc sailFunction) bool {
	for navS := range inputCh(navFile) {
		sailFunc(navS)
	}
	return true
}

type sailFunction func(navS string) (int, int)

func (this *Vessel) sail(navS string) (int, int) {
	parsed := this.reg_course.FindStringSubmatch(navS)
	dir := parsed[1]
	dist, _ := strconv.Atoi(parsed[2])
	if "R" == dir || "L" == dir {
		dist = dist / 90
		if "R" == dir {
			this.currFacing = (this.currFacing + dist) % 4
		} else {
			this.currFacing = (this.currFacing + 40 - dist) % 4
		}
		//fmt.Printf("   -> '%s' -> (%d,%d) turn %s by %d (facings[%d]: %s)\n", navS, this.Pos_x, this.Pos_y, dir, dist, this.currFacing, this.facings[this.currFacing])
		return this.Pos_x, this.Pos_y
	}
	if "F" == dir {
		dir = this.facings[this.currFacing]
	}
	if "N" == dir {
		this.Pos_y += dist
	} else if "E" == dir {
		this.Pos_x += dist
	} else if "S" == dir {
		this.Pos_y -= dist
	} else if "W" == dir {
		this.Pos_x -= dist
	}
	//fmt.Printf("   -> '%s' -> (%d,%d) -> %s by %d (facing: %s)\n", navS, this.Pos_x, this.Pos_y, dir, dist, this.facings[this.currFacing])
	return this.Pos_x, this.Pos_y
}

func (this *Vessel) sailToWaypoint(navS string) (int, int) {
	parsed := this.reg_course.FindStringSubmatch(navS)
	dir := parsed[1]
	dist, _ := strconv.Atoi(parsed[2])
	if "R" == dir || "L" == dir {
		dist = dist / 90
		for i := 0; i < dist; i++ {
			if "R" == dir {
				this.Waypoint_x, this.Waypoint_y = this.Waypoint_y, -this.Waypoint_x
			} else {
				this.Waypoint_x, this.Waypoint_y = -this.Waypoint_y, this.Waypoint_x
			}
		}
		//fmt.Printf("   -> '%s' curr-{p:(%d,%d), w:(%d,%d)} rotated waypoint %s by %d\n", navS, this.Pos_x, this.Pos_y, this.Waypoint_x, this.Waypoint_y, dir, dist)
		return this.Pos_x, this.Pos_y
	}
	if "F" == dir {
		this.Pos_x += this.Waypoint_x * dist
		this.Pos_y += this.Waypoint_y * dist
		//fmt.Printf("   -> '%s' curr-{p:(%d,%d), w:(%d,%d)} sailed to waypoint by %d\n", navS, this.Pos_x, this.Pos_y, this.Waypoint_x, this.Waypoint_y, dist)
		return this.Pos_x, this.Pos_y
	}
	if "N" == dir {
		this.Waypoint_y += dist
	} else if "E" == dir {
		this.Waypoint_x += dist
	} else if "S" == dir {
		this.Waypoint_y -= dist
	} else if "W" == dir {
		this.Waypoint_x -= dist
	}
	//fmt.Printf("   -> '%s' curr-{p:(%d,%d), w:(%d,%d)} moved waipoint to %s by %d\n", navS, this.Pos_x, this.Pos_y, this.Waypoint_x, this.Waypoint_y, dir, dist)
	return this.Pos_x, this.Pos_y
}
