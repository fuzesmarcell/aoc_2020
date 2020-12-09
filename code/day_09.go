package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
	"math"
)

func main() {
	data, err := ioutil.ReadFile("../inputs/day_09.txt")
	if err == nil {
		split_data := strings.Split(string(data), "\r\n")
		encrypted_numbers := make([]int, len(split_data))

		for i, string_number := range split_data {
			encrypted_numbers[i], _ = strconv.Atoi(string_number)
		}

		invalid_number := 0
		preamble_length := 25
		for i := preamble_length; i < len(encrypted_numbers); i++ {
			previous_numbers := encrypted_numbers[i - preamble_length:i]

			has_sum_of_two := false
			for j, a := range previous_numbers {
				for k, b := range previous_numbers {
					if j != k {
						if a + b == encrypted_numbers[i] {
							has_sum_of_two = true
							break
						}
					}
				}
				if has_sum_of_two {
					break
				}
			}
			
			if !has_sum_of_two {
				invalid_number = encrypted_numbers[i]
				fmt.Println("Part 01:", invalid_number)
				break
			}
		}

		for k := 0; k < len(encrypted_numbers); k++ {
			for i := 1; i < len(encrypted_numbers) - k; i++ {
				sum := 0
				min := math.MaxInt64
				max := math.MinInt64
				for j := 0; j < i; j++ {
					value := encrypted_numbers[k + j]
					if value < min {
						min = value
					}
					if value > max {
						max = value
					}

					sum += value
				}

				if sum > invalid_number {
					break
				}
				
				if sum == invalid_number {
					fmt.Println("Result 02:", min + max)
					return
				}
			}
		}
	}
}