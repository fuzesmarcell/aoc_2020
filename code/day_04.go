package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

const (
	byr string = "byr"
	iyr string = "iyr"
	eyr string = "eyr"
	hgt string = "hgt"
	hcl string = "hcl"
	ecl string = "ecl"
	pid string = "pid"
	cid string = "cid"
)

func is_numerical_field_in_range(field string, min int, max int) bool {

	result := false
	re := regexp.MustCompile(`\d+`)
	found_string := re.FindString(field)
	if found_string != "" {
		number, _ := strconv.Atoi(found_string)
		if number >= min && number <= max {
			result = true								
		}
	}

	return result
}

func main() {
	
	data, err := ioutil.ReadFile("../inputs/day_04.txt")
	
	fields_that_should_be_present := []string {
		byr,
		iyr,
		eyr,
		hgt,
		hcl,
		ecl,
		pid,
	}

	valid_passport_and_fields := 0
	valid_passports := 0
	if err == nil {
		split_data := strings.Split(string(data), "\r\n\r\n")
		fmt.Println("Number of Passports:", len(split_data))
		for _, passport := range split_data {
			is_valid := true
			for _, field_string := range fields_that_should_be_present {
				is_valid = strings.Contains(passport, field_string)
				if !is_valid {
					break
				}
			}

			if is_valid {
				all_fields_valid := true
				valid_passports += 1

				stripped_passport := strings.ReplaceAll(passport, "\r\n", " ")
				// NOTE(fuzes): This passport contains all the required fields time to parse them
				string_fields := strings.Split(stripped_passport, " ")
				for _, field := range string_fields {
					split_field := strings.Split(field, ":")

					field_name := split_field[0]
					field_value := split_field[1]

					is_field_valid := false
					switch field_name {
					case byr:
						is_field_valid = is_numerical_field_in_range(field_value, 1920, 2002)
					case iyr:
						is_field_valid = is_numerical_field_in_range(field_value, 2010, 2020)
					case eyr:
						is_field_valid = is_numerical_field_in_range(field_value, 2020, 2030)
					case hgt:
						if strings.Contains(field_value, "cm") {
							is_field_valid = is_numerical_field_in_range(field_value, 150, 193)
						} else if strings.Contains(field_value, "in") {
							is_field_valid = is_numerical_field_in_range(field_value, 59, 76)
						}
					case hcl:
						re := regexp.MustCompile(`#([0-9a-fA-F]+)`)
						found_string := re.FindAllStringSubmatch(field, -1)
						if len(found_string) > 0 {
							is_field_valid = len(found_string[0][1]) == 6
						}
					case ecl:
						var valid_ecl_colors = [] string {
							"amb",
							"blu",
							"brn",
							"gry",
							"grn",
							"hzl",
							"oth",
						}

						for _, ecl_color := range valid_ecl_colors {
							if field_value == ecl_color {
								is_field_valid = true
								break
							}
						}
					case pid:
						re := regexp.MustCompile(`\d+`)
						found_string := re.FindString(field)
						if found_string != "" {
							is_field_valid = len(found_string) == 9
						}
					case cid:
						is_field_valid = true
					}

					fmt.Printf("{%s, %s, %t}, ", field_name, field_value, is_field_valid)
					if !is_field_valid {
						all_fields_valid = false
						break
					}
				}

				if all_fields_valid {
					valid_passport_and_fields += 1
				}

				fmt.Printf("-> Fields are all valid: %t", all_fields_valid)
				fmt.Printf("\n")
			}
		}
	}

	fmt.Println("Part 01 valid passports:", valid_passports)
	fmt.Println("Part 02 valid passports:", valid_passport_and_fields)
}