package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	token_type_EOF          int = 0
	token_type_number       int = 1
	token_type_open_paren   int = 2
	token_type_closed_paren int = 3
	token_type_asterisk     int = 4
	token_type_plus         int = 5
	token_type_pipe         int = 6
	token_type_colon        int = 7
	token_type_identifier   int = 8
	token_type_string       int = 9
)

type token struct {
	kind        int
	text        string
	value       int64
	string_text string
}

type parser struct {
	is_done     bool
	token_index int
	tokens      []string
}

func GetTokenForIndex(p *parser, i int) token {
	var result token
	if i >= len(p.tokens) {
		p.is_done = true
		result.kind = token_type_EOF
		result.text = ""
		result.value = 0

		return result
	}

	text := p.tokens[i]

	result.text = text
	result.value = 0

	switch text[0] {
	case ':':
		result.kind = token_type_colon
	case '|':
		result.kind = token_type_pipe
	case '"':
		result.kind = token_type_string
		result.string_text = strings.Trim(result.text, "\"")
	case '*':
		result.kind = token_type_asterisk
	case '+':
		result.kind = token_type_plus
	case '(':
		result.kind = token_type_open_paren
	case ')':
		result.kind = token_type_closed_paren
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		result.kind = token_type_number
		result.value, _ = strconv.ParseInt(result.text, 10, 64)
	}

	return result
}

func GetNextToken(p *parser) token {

	result := GetTokenForIndex(p, p.token_index)
	p.token_index += 1

	return result
}

func ExpectTokenBool(p *parser, token_type int) bool {
	token := GetNextToken(p)
	if token.kind == token_type {
		return true
	}

	error_msg := fmt.Sprintf("Error at index %d", p.token_index)
	panic(error_msg)
}

func ExpectToken(p *parser, token_type int) token {
	token := GetNextToken(p)
	if token.kind == token_type {
		return token
	}

	error_msg := fmt.Sprintf("Error at index %d", p.token_index)
	panic(error_msg)
}

func PeekNextToken(p *parser) token {
	result := GetTokenForIndex(p, p.token_index)
	return result
}

type satellite_rule struct {
	is_literal   bool
	literal      byte
	sub_rules    [][]int64
	permutations []string
}

func EvaluateAllStringPermutations(rule_table map[int64]satellite_rule, rule_number int64) (bool, byte) {
	rule := rule_table[rule_number]

	if rule.is_literal {
		return true, rule.literal
	} else {
		for _, sub_rule := range rule.sub_rules {
			fmt.Println(sub_rule)
			var permutation_list []byte
			for _, n := range sub_rule {
				is_lit, lit_char := EvaluateAllStringPermutations(rule_table, n)
				if is_lit {
					permutation_list = append(permutation_list, lit_char)
				}
				rule.permutations = append(rule.permutations, string(permutation_list))
			}

		}
	}

	rule_table[rule_number] = rule
	return false, ' '
}

func main() {
	data, _ := ioutil.ReadFile("../../inputs/day_19.txt")

	rule_and_test_data := strings.Split(string(data), "\r\n\r\n")

	rules := rule_and_test_data[0]
	var rule_table map[int64]satellite_rule
	rule_table = make(map[int64]satellite_rule)

	for _, rule_line := range strings.Split(rules, "\r\n") {
		fmt.Println(rule_line)
		var p parser
		p.is_done = false
		p.token_index = 0
		p.tokens = strings.Split(rule_line, " ")

		rule_number_token := ExpectToken(&p, token_type_number)
		if ExpectTokenBool(&p, token_type_colon) {
			next_token := GetNextToken(&p)
			var r satellite_rule
			if next_token.kind == token_type_number {
				r.is_literal = false
				for true {
					var rule_array []int64
					for next_token.kind == token_type_number {
						rule_array = append(rule_array, next_token.value)
						next_token = GetNextToken(&p)
					}

					fmt.Println(rule_array)
					r.sub_rules = append(r.sub_rules, rule_array)

					if next_token.kind == token_type_EOF {
						break
					} else {
						if next_token.kind != token_type_pipe {
							panic("Expected pipe to seperate rule.")
						} else {
							next_token = ExpectToken(&p, token_type_number)
						}
					}
				}

			} else if next_token.kind == token_type_string {
				r.is_literal = true
				r.literal = byte(next_token.string_text[0])
			}

			rule_table[rule_number_token.value] = r
		}
	}

	EvaluateAllStringPermutations(rule_table, 0)
	for key, value := range rule_table {
		fmt.Println(key, value)
	}
}
