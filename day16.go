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
	return tc.MyTicketValues("departure")
}

type TicketComputer struct {
	fieldNames   []string
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

func (this *TicketComputer) AllFieldNames() []string {
	fieldNames := make([]string, 0)
	for _, validator := range this.validators {
		fieldNames = append(fieldNames, validator.name)
	}
	return fieldNames
}

func (this *TicketComputer) FieldNames() []string {
	allFieldNames := this.AllFieldNames()
	fieldNames := make([][]string, len(allFieldNames))
	// defining possible names by removing impossible names
	for i, _ := range allFieldNames {
		fieldNames[i] = make([]string, len(allFieldNames))
		copy(fieldNames[i], allFieldNames)
		for _, ticket := range this.validTickets {
			for _, impName := range ticket.fields[i].impNames {
				fieldNames[i] = remove(fieldNames[i], impName)
			}
		}
	}

	//reducing possible names by names occuring only once
	removed := make([]string, 0)
	for true {
		wasChange := false
		for i, names := range fieldNames {
			if len(names) == 1 && !contains(removed, names[0]) {
				for j, _ := range fieldNames {
					if i == j {
						continue
					}
					fieldNames[j] = remove(fieldNames[j], names[0])
				}
				removed = append(removed, names[0])
				wasChange = true
			}
		}
		if !wasChange {
			break
		}
	}

	//defining final field names
	this.fieldNames = make([]string, 0)
	for _, names := range fieldNames {
		this.fieldNames = append(this.fieldNames, names[0])
	}
	return this.fieldNames
}

func (this *TicketComputer) MyTicketValues(field string) int {
	if len(this.fieldNames) == 0 {
		this.FieldNames()
	}
	valsMul := 1
	for i, fieldName := range this.fieldNames {
		if strings.Contains(fieldName, field) {
			valsMul *= this.myTicket.fields[i].no
		}
	}
	return valsMul
}

type Ticket struct {
	fields     []TicketField
	validators []TicketFieldValidator
}

type TicketField struct {
	impNames []string
	no       int
	valid    bool
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
	field.valid = false
	this.fields = append(this.fields, *field)
}

func (this *Ticket) IsValid() bool {
	return len(FilterTicketFields(this.fields, func(field TicketField) bool { return !field.valid })) == 0
}

func (this *Ticket) ErrRate() int {
	errRate := 0
	for _, field := range this.fields {
		if !field.valid {
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

func (this *TicketFieldValidator) ValidateTicket(ticket *Ticket) {
	for i, field := range ticket.fields {
		if this.validateField(field) {
			ticket.fields[i].valid = true
		} else {
			ticket.fields[i].impNames = append(ticket.fields[i].impNames, this.name)
		}
	}
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
