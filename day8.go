package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


func Day8_1(filename string) int {
	c := NewComputer(filename)
	c.Run()
	return c.Acc
}

func Day8_2(filename string) int {
	acc := 0
	return acc
}

type computer struct {
	Program []computerCommand
	currPos int
	nextPos int
	Acc int
	State string
}

func NewComputer(programFile string) computer {
	c := new(computer)
	c.Program = make([]computerCommand, 0)
	c.currPos = 0
	c.nextPos = 0
	c.Acc = 0
	c.State = ""

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
	return *c
}

func (c* computer) Run() {
	//fmt.Println("=== BEGIN ===")
	for c.runNextCommand() {
		//fmt.Printf("c.Program[%d].count: %d\n", c.currPos, c.Program[c.currPos].count)
	}
	fmt.Printf("=== %s ===\n", c.State)
}

func (c* computer) runNextCommand() bool {
	if c.nextPos == len(c.Program) {
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
		//fmt.Printf("   [%d] ACC: %d\n", c.currPos, command.args[0])
	} else if "jmp" == command.cmd {
		c.nextPos += command.args[0]
		//fmt.Printf("   [%d] JMP: %d -> %d\n", c.currPos, command.args[0], c.nextPos)
	} else  if "nop" == command.cmd {
		//fmt.Printf("   [%d] NOP\n", c.currPos)
	}
	if c.nextPos == c.currPos {
		c.nextPos++
	}
	command.count++
	return true 
}

type computerCommand struct {
	cmd string
	args []int
	count int
}

