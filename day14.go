package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day14_1(filename string) int {
	fmt.Printf("")
	comp := NewDockingComputer(filename)
	comp.Run()
	return int(comp.SumMem())
}

func Day14_2(filename string) int {
	fmt.Printf("")
	comp := NewDockingComputer(filename)
	fmt.Println("comp: ", comp)
	return 0
}

type DockingComputer struct {
	programs []DockingProgram
	mem      [100000]int
}

func NewDockingComputer(filename string) DockingComputer {
	comp := new(DockingComputer)
	comp.programs = make([]DockingProgram, 0)
	prog := new(DockingProgram)
	for fline := range inputCh(filename) {
		cmdArg := strings.Split(fline, " = ")
		if cmdArg[0] == "mask" {
			if len(prog.mask) > 0 {
				comp.programs = append(comp.programs, *prog)
			}
			prog = new(DockingProgram)
			prog.mask = cmdArg[1]
			prog.parseMask()
		} else {
			cmd := new(DockingCommand)
			cmd.ind, _ = strconv.Atoi(string(cmdArg[0][4 : len(cmdArg[0])-1]))
			val, _ := strconv.Atoi(cmdArg[1])
			cmd.val = int(val)
			prog.commands = append(prog.commands, *cmd)
		}
	}
	comp.programs = append(comp.programs, *prog)
	return *comp
}

func (this *DockingComputer) Run() {
	for _, prog := range this.programs {
		prog.Run(&this.mem)
	}
}

func (this *DockingComputer) SumMem() int {
	sum := 0
	for _, memVal := range this.mem {
		sum += memVal
	}
	return sum
}

func (this *DockingComputer) PrintMem() string {
	str := ""
	for _, memVal := range this.mem {
		if memVal != 0 {
			str += fmt.Sprintf("%d, ", memVal)
		}
	}
	return str

}

type DockingProgram struct {
	mask     string
	maskVals []int
	maskType []bool
	commands []DockingCommand
}

type DockingCommand struct {
	ind int
	val int
}

func (this *DockingProgram) parseMask() {
	this.maskVals = make([]int, 0)
	this.maskType = make([]bool, 0)
	pow2 := 1
	for i := len(this.mask) - 1; i >= 0; i-- {
		if "X" != string(this.mask[i]) {
			this.maskVals = append(this.maskVals, int(pow2))
			this.maskType = append(this.maskType, "1" == string(this.mask[i]))
		}
		pow2 *= 2
	}
}

func (this *DockingProgram) Run(mem *[100000]int) {
	for _, cmd := range this.commands {
		newVal := cmd.val
		for i, maskVal := range this.maskVals {
			byteEnabled := newVal&maskVal > 0
			if this.maskType[i] && !byteEnabled {
				newVal += maskVal
			} else if !this.maskType[i] && byteEnabled {
				newVal -= maskVal
			}
		}
		//fmt.Printf("setting mem[%d] := %d -> %d\n", cmd.ind, cmd.val, newVal)
		(*mem)[cmd.ind] = newVal
	}
}
