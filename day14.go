package main

import (
	"fmt"
	"strconv"
	"strings"
)

const MEM_SIZE = 100000000

func Day14_1(filename string) int {
	fmt.Printf("")
	comp := NewDockingComputer(filename)
	comp.Run(comp.RunNormalProgram)
	return int(comp.SumMem())
}

func Day14_2(filename string) int {
	fmt.Printf("")
	comp := NewDockingComputer(filename)
	comp.Run(comp.RunMemoryProgram)
	return int(comp.SumMem())
}

type DockingComputer struct {
	programs []DockingProgram
	mem      map[int]int
}

func NewDockingComputer(filename string) DockingComputer {
	comp := new(DockingComputer)
	comp.mem = make(map[int]int)
	comp.programs = make([]DockingProgram, 0)
	prog := new(DockingProgram)
	for fline := range inputCh(filename) {
		cmdArg := strings.Split(fline, " = ")
		if cmdArg[0] == "mask" {
			if len(prog.maskString) > 0 {
				comp.programs = append(comp.programs, *prog)
			}
			prog = new(DockingProgram)
			prog.maskString = cmdArg[1]
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

type RunFunc func(prog *DockingProgram, mem *map[int]int)

func (this *DockingComputer) Run(runFunction RunFunc) {
	for _, prog := range this.programs {
		runFunction(&prog, &this.mem)
	}
}

func (this *DockingComputer) RunNormalProgram(prog *DockingProgram, mem *map[int]int) {
	prog.RunNormalMode(&this.mem)
}

func (this *DockingComputer) RunMemoryProgram(prog *DockingProgram, mem *map[int]int) {
	prog.RunMemoryMode(&this.mem)
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
	maskString string
	masks      []DockingMask
	commands   []DockingCommand
}

type DockingMask struct {
	val int
	typ int
}

type DockingCommand struct {
	ind int
	val int
}

func (this *DockingProgram) parseMask() {
	this.masks = make([]DockingMask, 0)
	pow2 := 1
	for i := len(this.maskString) - 1; i >= 0; i-- {
		mask := new(DockingMask)
		mask.val = int(pow2)
		if "X" != string(this.maskString[i]) {
			mask.typ, _ = strconv.Atoi(string(this.maskString[i]))
		} else {
			mask.typ = -1
		}
		this.masks = append(this.masks, *mask)
		pow2 *= 2
	}
}

func (this *DockingProgram) RunNormalMode(mem *map[int]int) {
	for _, cmd := range this.commands {
		newVal := cmd.val
		for _, mask := range this.masks {
			byteEnabled := newVal&mask.val > 0
			if mask.typ == 1 && !byteEnabled {
				newVal += mask.val
			} else if mask.typ == 0 && byteEnabled {
				newVal -= mask.val
			}
		}
		//fmt.Printf("--- setting mem[%d] := %d -> %d\n", cmd.ind, cmd.val, newVal)
		(*mem)[cmd.ind] = newVal
	}
}

func (this *DockingProgram) RunMemoryMode(mem *map[int]int) {
	for _, cmd := range this.commands {
		inds := make([]int, 0)
		ind := cmd.ind
		for _, mask := range this.masks {
			byteEnabled := ind&mask.val > 0
			if mask.typ == 1 && !byteEnabled {
				ind += mask.val
			}
		}
		inds = append(inds, ind)
		//fmt.Printf("    + %d\n", ind)
		for _, mask := range this.masks {
			for _, ind := range inds {
				if mask.typ >= 0 || mask.val > ind {
					break
				}
				byteEnabled := ind&mask.val > 0
				if mask.typ == -1 && !byteEnabled {
					inds = append(inds, ind+mask.val)
					//fmt.Printf("    + %d\n", ind+mask.val)
				} else if mask.typ == -1 && byteEnabled {
					inds = append(inds, ind-mask.val)
					//fmt.Printf("    + %d\n", ind-mask.val)
				}
			}
		}
		//fmt.Printf("setting cmd.ind: %d => %s => mem[%d] := %d\n", cmd.ind, this.maskString, inds, cmd.val)
		for _, ind := range inds {
			(*mem)[ind] = cmd.val
		}
		//fmt.Printf(" added %d, total %d -- %d\n", len(inds), len(*mem), *mem)
		//fmt.Printf(" added %d, total %d\n", len(inds), len(*mem))
	}
}
