package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("../inputs/day_06.txt")

	if err == nil {
		puzzle_input := strings.Split(string(data), "\r\n\r\n")

		sum_of_counts := 0
		sum_of_counts_part_02 := 0
		for _, group_answers := range puzzle_input {
			split_answers := strings.Split(group_answers, "\r\n")
			
			var answers_table map[rune]bool
			answers_table = make(map[rune]bool)

			var hit_answer_table map[rune]int
			hit_answer_table = make(map[rune]int)

			for _, answer := range split_answers {
				for _, char := range answer {
					answers_table[char] = true
					if hit_answer_table[char] == 0 {
						hit_answer_table[char] = 1
					} else {
						hit_answer_table[char] += 1
					}
				}
			}

			sum_of_counts += len(answers_table)
			// fmt.Println("-----", len(answers_table), answers_table)

			number_of_answers := len(split_answers)
			for _, value := range hit_answer_table {
				if value == number_of_answers {
					sum_of_counts_part_02 += 1
				}
			}

			// fmt.Println(len(split_answers), hit_answer_table, sum_of_counts_part_02)
		}

		fmt.Println("Part 01 Sum of counts:", sum_of_counts)
		fmt.Println("Part 02 Sum of counts:", sum_of_counts_part_02)
	}
}