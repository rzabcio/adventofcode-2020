package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


func Day8_1(filename string) int {
	fmt.Printf("")
	c := NewComputer(filename)
	c.Run()
	return c.Acc
}

func Day8_2(filename string) int {
	c := NewComputer(filename)
	c.RunSelfRepair()
	return c.Acc
}

type computer struct {
	origProgram []computerCommand
	Program []computerCommand
	currPos int
	nextPos int
	Acc int
	State string
}

func NewComputer(programFile string) computer {
	c := new(computer)
	c.Program = make([]computerCommand, 0)

	reg_line := regexp.MustCompile(`^(.*) (.\d*)`)
	for line := range inputCh(programFile) {
		parsed := reg_line.FindStringSubmatch(line)
		command := new(computerCommand)
		command.cmd = parsed[1]
		command.args = make([]int, 0)
		for _, argS := range(strings.Split(parsed[2], ",")) {
			arg, _ := strconv.Atoi(argS)
			command.args = append(command.args, arg) 
		}
		c.Program = append(c.Program, *command)
	}
	c.origProgram = make([]computerCommand, len(c.Program))
	copy(c.origProgram, c.Program)
	c.Reset()
	return *c
}

func (c* computer) Reset() {
	copy(c.Program, c.origProgram)
	c.currPos = 0
	c.nextPos = 0
	c.Acc = 0
	c.State = ""
}

func (c* computer) Run() bool {
	for c.runNextCommand() {
	}
	return strings.Index(c.State, "OK") > -1  //successfull run sets "OK"
}

func (c* computer) runNextCommand() bool {
	if c.nextPos == len(c.Program) {  // EOF - successfull
		c.State = "OK"
		return false
	}
	command := &c.Program[c.nextPos]
	if command.count == 1 {
		c.State = "ERR: infinite loop"
		return false
	}
	c.currPos = c.nextPos
	if "acc" == command.cmd {
		c.Acc += command.args[0]
	} else if "jmp" == command.cmd {
		c.nextPos += command.args[0]
	}
	if c.nextPos == c.currPos {
		c.nextPos++
	}
	command.count++
	return true 
}

func (c* computer) RunSelfRepair() bool {
	for commandPos, command := range c.origProgram {
		c.Reset()
		if command.cmd == "jmp" {
			c.Program[commandPos].cmd = "nop"
		} else if command.cmd == "nop" {
			c.Program[commandPos].cmd = "jmp"
		} else {
			continue
		}
		if c.Run() {
			return true
		}
	}
	return false
}

type computerCommand struct {
	cmd string
	args []int
	count int
}

