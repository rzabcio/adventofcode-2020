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
	return tc.ValidateTickets()
}

func Day16_2(filename string) int {
	tc := NewTicketComputer(filename)
	tc.ValidateTickets()
	tc.FieldNames()
	return 0
}

type TicketComputer struct {
	validators   []TicketFieldValidator
	myTicket     Ticket
	nearTickets  []Ticket
	validTickets []Ticket
}

func NewTicketComputer(programFile string) TicketComputer {
	this := new(TicketComputer)
	this.validators = make([]TicketFieldValidator, 0)
	this.nearTickets = make([]Ticket, 0)
	reg_validator := regexp.MustCompile(`^(.*): (\d*)-(\d*) or (\d*)-(\d*)`)
	for line := range inputCh(programFile) {
		parsed := reg_validator.FindStringSubmatch(line)
		if len(parsed) > 0 {
			validator := new(TicketFieldValidator)
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
		newTicket := NewTicket()
		for _, noStr := range ticketArr {
			no, _ := strconv.Atoi(noStr)
			newTicket.AddField(no)
		}
		if len(this.myTicket.fields) == 0 {
			this.myTicket = newTicket
		} else {
			this.nearTickets = append(this.nearTickets, newTicket)
		}
	}
	return *this
}

func (this *TicketComputer) ValidateTickets() int {
	errRateSum := 0
	for _, ticket := range this.nearTickets {
		for _, validator := range this.validators {
			validator.ValidateTicket(&ticket)
		}
		if ticket.ErrRate() == 0 {
			this.validTickets = append(this.validTickets, ticket)
		}
		errRateSum += ticket.ErrRate()
	}
	return errRateSum
}

func (this *TicketComputer) FieldNames() []string {
	names := make([]string, 0)
	for i, _ := range this.validTickets[0].fields {
		//for _, ticket := range this.validTickets {
		for _, ticket := range this.nearTickets {
			fmt.Printf(" %d, ", len(ticket.fields[i].possNames))
			if len(ticket.fields[i].possNames) == 1 {
				names = append(names, ticket.fields[i].possNames[0])
			}
		}
		fmt.Printf("\n")
		if len(names) < i+1 {
			names = append(names, "?")
		}
	}
	fmt.Printf("%s\n", names)
	return names
}

type Ticket struct {
	fields     []TicketField
	validators []TicketFieldValidator
}

type TicketField struct {
	possNames []string
	name      string
	no        int
	validated bool
}

func NewTicket() Ticket {
	this := new(Ticket)
	this.fields = make([]TicketField, 0)
	this.validators = make([]TicketFieldValidator, 0)
	return *this
}

func (this *Ticket) AddField(no int) {
	field := new(TicketField)
	field.no = no
	field.possNames = make([]string, 0)
	this.fields = append(this.fields, *field)
}

func (this *Ticket) IsValid() bool {
	return len(FilterTicketFields(this.fields, func(field TicketField) bool { return !field.validated })) == 0
}

func (this *Ticket) ErrRate() int {
	errRate := 0
	for _, field := range this.fields {
		if !field.validated {
			errRate += field.no
		}
	}
	return errRate
}

type TicketFieldValidator struct {
	name string
	min  []int
	max  []int
}

func (this *TicketFieldValidator) ValidateTicket(ticket *Ticket) bool {
	valid := false
	for i, field := range ticket.fields {
		if this.validateField(field) && len(field.name) == 0 {
			//ticket.fields[i].name = this.name
			ticket.fields[i].possNames = append(ticket.fields[i].possNames, this.name)
			ticket.fields[i].validated = true
			continue
		} else {
			valid = valid || true
		}
	}
	return valid
}

func (this *TicketFieldValidator) validateField(field TicketField) bool {
	result := false
	for i, _ := range this.min {
		result = result || (this.min[i] <= field.no && field.no <= this.max[i])
	}
	//fmt.Printf("validating if %d is in %d: %t\n", field.no, this, result)
	return result
}

// TOOLS ////////////////////////////////////////////////////////////
func FilterTickets(arr []Ticket, cond func(Ticket) bool) []Ticket {
	result := []Ticket{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func FilterTicketFields(arr []TicketField, cond func(TicketField) bool) []TicketField {
	result := []TicketField{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
