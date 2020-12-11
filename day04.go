package main

import (
	"regexp"
	"strconv"
	"strings"
)


func Day4_1(filename string) int {
	validCount := 0
	for _, p := range parsePassports(filename) {
		// fmt.Println("--- passport: ", p)
		if validatePassport(p, false) {
			validCount++
		}
	}
	return validCount
}

func Day4_2(filename string) int {
	validCount := 0
	for _, p := range parsePassports(filename) {
		if validatePassport(p, true) {
			validCount++
		}
	}
	return validCount
}

func parsePassports(filename string) []map[string]string {
	passports := make([]map[string]string, 0)
	p := make(map[string]string)
	for line := range inputCh(filename) {
		if line == "" {
			passports = append(passports, p)
			p = make(map[string]string)
			continue
		}
		parsePassportLine(p, line)
	}
	passports = append(passports, p)
	return passports
}
func parsePassportLine(p map[string]string, line string) {
	for _, entry := range strings.Split(line, " ") {
		split := strings.Split(entry, ":")
		key, val := split[0], split[1]
		p[key] = val
	}
}
func validatePassport(p map[string]string, validateFields bool) bool {
	//fmt.Println("--- passport: ", p)
	_, hasCid := p["cid"]
	if len(p) < 8 && (len(p)!=7 || hasCid) {
		//fmt.Println("    NOT OK: fields")
		return false
	}
	if !validateFields {
		return true
	}
	for key, val := range p {
		v_reg := v_regs[key]
		if v_reg != nil && !v_reg.MatchString(val) {
				//fmt.Println("    NOT OK:  v_reg: ", key)
				return false
		}
		v_func := v_funcs[key]
		if v_func != nil && !v_func(val) {
				//fmt.Println("    NOT OK: v_func: ", key)
				return false
		}
	}
	return true
}
var v_regs = map[string]*regexp.Regexp {
	"byr": regexp.MustCompile(`^\d{4}$`),
	"iyr": regexp.MustCompile(`^\d{4}$`),
	"eyr": regexp.MustCompile(`^\d{4}$`),
	"hgt": regexp.MustCompile(`^(\d{2,4})(cm|in)$`),
	"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
	"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
	"pid": regexp.MustCompile(`^\d{9}$`),
}
var v_funcs = map[string]func(string) bool{
	"byr": func(s string) bool {
					i, _ := strconv.Atoi(s)
					return 1920<=i && i<=2002
				},
	"iyr": func(s string) bool {
					i, _ := strconv.Atoi(s)
					return 2010<=i && i<=2020
				},
	"eyr": func(s string) bool {
					i, _ := strconv.Atoi(s)
					return 2020<=i && i<=2030
				},
	"hgt": func(s string) bool {
					heightUnit := v_regs["hgt"].FindStringSubmatch(s)
					height, _ := strconv.Atoi(heightUnit[1])
					if heightUnit[2] == "in" {
						return 59<=height && height<=76
					}
					return 150<=height && height<=193
				},
}

