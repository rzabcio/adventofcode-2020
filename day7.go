package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


func Day7_1(filename string) int {
	bagRules := readBagRules(filename)
	containers := findContainers(bagRules, []string{"shiny gold"})
	fmt.Println("containers: ", containers)
	return len(containers)-1
}

func Day7_2(filename string) int {
	bagRules := readBagRules(filename)
	return calcContainerCost(bagRules, "shiny gold")
}

func readBagRules(filename string) map[string][]string {
	bagRules := make(map[string][]string, 0)
	r_line := regexp.MustCompile(`^(.*) bags contain (.*)`)
	r_bag := regexp.MustCompile(`(\d) (.*) bag.*$`)

	for line := range inputCh(filename) {
		parsed := r_line.FindStringSubmatch(line)
		container := parsed[1]
		bagRules[container] = make([]string, 0)
		for _, bagS := range(strings.Split(strings.ReplaceAll(parsed[2], "no other", "0"), ",")) {
			bag := r_bag.FindStringSubmatch(bagS)
			if len(bag) == 0 {
				continue
			}
			bagCount, _ := strconv.Atoi(bag[1])
			for i := 0; i<bagCount; i++ { // add entry for every needed luggage (to count it in the future)
				bagRules[container] = append(bagRules[container], bag[2])
			}
		}
	}
	return bagRules
}

func findContainers(bagRules map[string][]string, containers []string) []string {
	// immediate level - all subcontainers
	for cont, bags := range(bagRules) {
		if contains(bags, containers[0]) {
			containers = append(containers, cont)
		}
	}
	// first sub-container level (recurrency)
	if len(containers) > 1 {
		for _, subCont := range findContainers(bagRules, containers[1:]) {
			if !contains(containers, subCont) {
				containers = append(containers, subCont)
			}
		}
	}
	return containers
}

func calcContainerCost(bagRules map[string][]string, container string) int {
	cost := 0
	if len(bagRules[container]) == 0 { // no sub-containers (lowest level)
		return 0
	}
	for _, subCont := range bagRules[container] {
		cost++																				// the cost of this subcontainter
		cost += calcContainerCost(bagRules, subCont)	// the cost of sub-subcontainers
	}
	return cost
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
