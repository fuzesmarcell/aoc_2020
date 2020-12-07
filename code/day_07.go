package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type rule_pair struct {
	amount int
	name string
}

type bag_rule struct {
	name string
	rule_set_len int
	rule_pairs []rule_pair
}

type hash_rule_table map[string]bag_rule

func RecursivlySearchForShinyGoldBag(rule_table hash_rule_table, rule bag_rule) bool {
	result := false
	for _, pair := range rule.rule_pairs {
		if pair.name == "shiny gold" {
			result = true
			break
		} else {
			new_rule := rule_table[pair.name]
			result = RecursivlySearchForShinyGoldBag(rule_table, new_rule)
			if result {
				break
			}
		}
	}

	return result
}

func RecursivlyFindAmountOfBagsForRule(rule_table hash_rule_table, rule bag_rule) int {
	result := 0
	for _, pair := range rule.rule_pairs {
		result += pair.amount
	}

	for _, pair := range rule.rule_pairs {
		amount_of_bags := RecursivlyFindAmountOfBagsForRule(rule_table, rule_table[pair.name])
		result += (amount_of_bags * pair.amount)
	}

	return result
}

func main() {

	data, err := ioutil.ReadFile("../inputs/day_07.txt")

	if err == nil {
		all_rules := strings.Split(string(data), "\r\n")

		var rule_table hash_rule_table
		rule_table = make(hash_rule_table)

		for _, rule := range all_rules {
			var new_rule bag_rule
			if strings.Contains(rule, "no other bags") {
				rule_ruleset_split := strings.Split(rule, "contain")
				rule_name_split := strings.Split(rule_ruleset_split[0], " ")
				rule_name := fmt.Sprintf("%s %s", rule_name_split[0], rule_name_split[1])

				new_rule.name = rule_name
				new_rule.rule_set_len = 0

				rule_table[new_rule.name] = new_rule

			} else {
				rule_ruleset_split := strings.Split(rule, "contain")
				rule_name_split := strings.Split(rule_ruleset_split[0], " ")
				rule_name := fmt.Sprintf("%s %s", rule_name_split[0], rule_name_split[1])
				
				rule_set := strings.Split(rule_ruleset_split[1], ",")

				new_rule.name = rule_name
				new_rule.rule_set_len = len(rule_set)
				new_rule.rule_pairs = make([]rule_pair, new_rule.rule_set_len)

				for i, sub_rule := range rule_set {
					sub_rule_split := strings.Split(sub_rule, " ")
					amount, _ := strconv.Atoi(sub_rule_split[1])
					sub_rule_name := fmt.Sprintf("%s %s", sub_rule_split[2], sub_rule_split[3])
					new_rule.rule_pairs[i] = rule_pair{amount, sub_rule_name}
				}

				rule_table[new_rule.name] = new_rule
			}

			// fmt.Println(new_rule)
		}

		can_contain_shiny_gold_bag_count := 0
		for _, rule := range rule_table {
			can_contain_shiny_gold_bag := RecursivlySearchForShinyGoldBag(rule_table, rule)
			if can_contain_shiny_gold_bag {
				can_contain_shiny_gold_bag_count += 1
			}

			// fmt.Println(rule, can_contain_shiny_gold_bag)
		}

		amount_of_bags := RecursivlyFindAmountOfBagsForRule(rule_table, rule_table["shiny gold"])

		fmt.Println("Part 01 solution:", can_contain_shiny_gold_bag_count)
		fmt.Println("Part 02 solution:", amount_of_bags)

	} else {
		fmt.Println("Could not oppen file!")
	}
}