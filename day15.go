package main

import (
	"fmt"
)

func Day15_1(filename string) int {
	fmt.Printf("")
	game := NewElvMemGame(filename)
	for 2020 >= game.NextTurn() {
	}
	return game.lastN
}

func Day15_2(filename string) int {
	fmt.Printf("")
	game := NewElvMemGame(filename)
	for 30000000 >= game.NextTurn() {
	}
	return game.lastN
}

type ElvMemGame struct {
	numbers map[int]*ElvMemGameNo
	lastN   int
	turn    int
}

func NewElvMemGame(filename string) ElvMemGame {
	this := new(ElvMemGame)
	this.numbers = make(map[int]*ElvMemGameNo, 0)
	this.turn = 1
	for n := range inputChInt(filename) {
		no := *new(ElvMemGameNo)
		no.no = n
		no.turns = make([]int, 0)
		(&no).addTurn(this.turn)
		this.numbers[n] = &no
		this.turn++
		this.lastN = n
	}
	return *this
}

func (this *ElvMemGame) NextTurn() int {
	//fmt.Printf("     t:%d, l:%d => %d", this.turn, this.lastN, this.numbers[this.lastN])
	no := this.numbers[this.lastN]
	this.lastN = no.age()
	if no, ok := this.numbers[this.lastN]; ok {
		no = this.numbers[this.lastN]
		no.addTurn(this.turn)
	} else {
		no := *new(ElvMemGameNo)
		no.no = this.lastN
		no.addTurn(this.turn)
		this.numbers[this.lastN] = &no
	}
	this.turn++
	return this.turn
}

type ElvMemGameNo struct {
	no    int
	turns []int
}

func (this *ElvMemGameNo) addTurn(turn int) {
	this.turns = append(this.turns, int(turn))
}

func (this *ElvMemGameNo) age() int {
	if len(this.turns) < 2 {
		return 0
	}
	return this.turns[len(this.turns)-1] - this.turns[len(this.turns)-2]
}

func (this *ElvMemGameNo) lastTurn() int {
	return this.turns[len(this.turns)-1]
}
