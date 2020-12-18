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
)

type token struct {
	kind  int
	text  string
	value int64
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

func PeekNextToken(p *parser) token {
	result := GetTokenForIndex(p, p.token_index)
	return result
}

func ParseFactor(p *parser) int64 {
	t := GetNextToken(p)
	if t.kind == token_type_number {
		return t.value
	} else if t.kind == token_type_open_paren {
		value := ParseExpression(p)
		next_token := GetNextToken(p)
		if next_token.kind != token_type_closed_paren {
			panic("Parantheses not closed")
		}
		return value
	}

	panic("Unexpected token in parse factor")
}

func ParseTerm(p *parser) int64 {
	value := ParseFactor(p)
	next_token := PeekNextToken(p)
	for next_token.kind == token_type_plus {
		// NOTE(fuzes): This is only needed so we advance the parser
		_ = GetNextToken(p)
		value += ParseFactor(p)

		next_token = PeekNextToken(p)
	}

	return value
}

func ParseExpression(p *parser) int64 {
	value := ParseTerm(p)
	next_token := PeekNextToken(p)
	for next_token.kind == token_type_asterisk {
		// NOTE(fuzes): This is only needed so we advance the parser
		_ = GetNextToken(p)
		value *= ParseTerm(p)
		next_token = PeekNextToken(p)
	}

	return value
}

func main() {
	data, _ := ioutil.ReadFile("../../inputs/day_18.txt")
	data_split := strings.Split(string(data), "\r\n")

	var sum int64
	for _, line := range data_split {
		split_line := strings.Split(line, " ")
		p := parser{false, 0, split_line}
		sum += ParseExpression(&p)
	}

	fmt.Println("Part 02:", sum)

}
