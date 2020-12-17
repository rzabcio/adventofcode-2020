package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day16_1(filename string) int {
	fmt.Printf("")
	tc := NewTicketComputer(filename)
	return len(tc.nearTickets)
}

func Day16_2(filename string) int {
	tc := NewTicketComputer(filename)
	return len(tc.nearTickets)
}

type TicketComputer struct {
	validators  []TicketValidator
	myTicket    []int
	nearTickets [][]int
}

func NewTicketComputer(programFile string) TicketComputer {
	this := new(TicketComputer)
	this.validators = make([]TicketValidator, 0)
	this.nearTickets = make([][]int, 0)
	reg_validator := regexp.MustCompile(`^(.*): (\d*)-(\d*) or (\d*)-(\d*)`)
	for line := range inputCh(programFile) {
		parsed := reg_validator.FindStringSubmatch(line)
		if len(parsed) > 0 {
			validator := new(TicketValidator)
			validator.name = parsed[1]
			validator.min = make([]int, (len(parsed)-2)/2)
			validator.max = make([]int, (len(parsed)-2)/2)
			for i := 2; i < len(parsed); i += 2 {
				validator.min[i/2-1], _ = strconv.Atoi(parsed[i])
				validator.max[i/2-1], _ = strconv.Atoi(parsed[i+1])
			}
			this.validators = append(this.validators, *validator)
			continue
		}

		ticketArr := strings.Split(line, ",")
		if len(ticketArr) < 2 {
			continue
		}
		newTicket := make([]int, 0)
		for _, noStr := range ticketArr {
			no, _ := strconv.Atoi(noStr)
			newTicket = append(newTicket, no)
		}
		if len(this.myTicket) == 0 {
			this.myTicket = newTicket
		} else {
			this.nearTickets = append(this.nearTickets, newTicket)
		}
	}
	return *this
}

func (this *TicketComputer) validateNearTickets() int {
	errRates := make([]int, len(this.validators)
	for i, validator := range this.validators {
		for _, no
	}
}

type TicketValidator struct {
	name string
	min  []int
	max  []int
}

func (this *TicketValidator) validate(no int) bool {
	result := false
	for i, _ := range this.min {
		result := result || (this.min <= no && no <= this.max)
	}
	return result
}
