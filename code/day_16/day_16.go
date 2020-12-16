package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

type ticket_rule struct {
	name string
	ranges [2][2]int
}

func IsFieldValid(rule ticket_rule, field int) bool {
	result := (field >= rule.ranges[0][0] && field <= rule.ranges[0][1]) || 
			  (field >= rule.ranges[1][0] && field <= rule.ranges[1][1])
	return result
}

func GetScanningErrorRateForSingleTicket(rules []ticket_rule, ticket []int) (bool, int) {

	result_count := 0
	result_bool := true

	for _, n := range ticket {
		not_valid_for_any_rule := true
		for _, rule := range rules {
			if IsFieldValid(rule, n) {
				not_valid_for_any_rule = false
				break
			}
		}

		if not_valid_for_any_rule {
			result_bool = false
			result_count += n
		}
	}

	return result_bool, result_count
}

func main() {
	data, _ := ioutil.ReadFile("../../inputs/day_16.txt")
	data_table := strings.Split(string(data), "\r\n\r\n")

	var ticket_rules []ticket_rule
	for _, rule_line := range strings.Split(data_table[0], "\r\n") {
		var ranges [2][2]int
		re := regexp.MustCompile(`(\d+)-(\d+)`)

		for i, r := range re.FindAllString(rule_line, 2) {
			ranges[i][0], _ = strconv.Atoi(strings.Split(r, "-")[0])
			ranges[i][1], _ = strconv.Atoi(strings.Split(r, "-")[1])
		}
		
		rule := ticket_rule{strings.Split(rule_line, ":")[0], ranges}

		ticket_rules = append(ticket_rules, rule)
	}

	fmt.Println(ticket_rules)

	var nearby_tickets [][]int
	for _, n_ticket_data := range strings.Split(data_table[2], "\r\n") {
		if n_ticket_data == "nearby tickets:" {
			continue
		}

		var arr []int
		for _, n := range strings.Split(n_ticket_data, ",") {
			number, _ := strconv.Atoi(n)
			arr = append(arr, number)
		}

		nearby_tickets = append(nearby_tickets, arr)
	}

	var valid_tickets [][]int
	sum := 0
	for _, ticket := range nearby_tickets {
		is_valid, count := GetScanningErrorRateForSingleTicket(ticket_rules, ticket)
		sum += count
		if is_valid {
			valid_tickets = append(valid_tickets, ticket)
		}
	}

	fmt.Println("Part 01:", sum)

	number_of_fields := len(valid_tickets[0])

	found_headers := make(map[string]bool)
	field_headers := make([]string, number_of_fields)

	for len(found_headers) < number_of_fields {
		for x := 0; x < number_of_fields; x++ {
			number_of_rules_valid := 0
			last_valid_rule_name := ""
			for _, rule := range ticket_rules {
				if _, in_map := found_headers[rule.name]; in_map {
					continue
				}
	
				all_fields_valid_for_rule := true
				for y := 0; y < len(valid_tickets); y++ {
					if !IsFieldValid(rule, valid_tickets[y][x]) {
						all_fields_valid_for_rule = false
						break
					}
				}
	
				if all_fields_valid_for_rule {
					number_of_rules_valid += 1
					last_valid_rule_name = rule.name
				}
			}
	
			if number_of_rules_valid == 1 {
				field_headers[x] = last_valid_rule_name
				found_headers[last_valid_rule_name] = true
			}
		}
	}

	result_02 := 1
	// NOTE(fuzes): A bit lazy to parse it we just copy pasted our ticket here... =)
	your_ticket := []int {151,139,53,71,191,107,61,109,157,131,67,73,59,79,113,167,137,163,149,127}
	for i, header_name := range field_headers {
		if strings.Contains(header_name, "departure") {
			result_02 *= your_ticket[i]
		}
	}

	fmt.Printf("\n")
	fmt.Println(result_02)
}