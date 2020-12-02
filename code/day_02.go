package main

import (
    "fmt"
    "io/ioutil"
    "text/scanner"
    "strings"
    "strconv"
)

type password_entry struct {
    min int
    max int
    char rune 
    password string
}

func is_entry_valid(entry password_entry) bool {
    
    count := 0
    for _, letter := range entry.password {
        if letter == entry.char {
            count += 1
        }
    }

    result := false
    if count >= entry.min && count <= entry.max {
        result = true
    }

    return result
}

func is_entry_valid_part_02(entry password_entry) bool {
    first_index := entry.min - 1
    second_index := entry.max - 1

    is_first_equal := rune(entry.password[first_index]) == entry.char
    is_second_equal := rune(entry.password[second_index]) == entry.char

    return is_first_equal != is_second_equal
}

func main() {

    fmt.Println("Saving christmas...")
    data, _ := ioutil.ReadFile("../inputs/day_02_part_01.txt")

    var parser scanner.Scanner
    parser.Init(strings.NewReader(string(data)))
    parser.Filename = "test"

    valid_password_counter_part_01 := 0
    valid_password_counter_part_02 := 0
    parsing := true
    for parsing {
        token := parser.Scan()
        if token != scanner.EOF {
            min := parser.TokenText()

            token = parser.Scan()
            if parser.TokenText() == "-" {
                
                token = parser.Scan()
                max := parser.TokenText()

                token = parser.Scan()
                char := parser.TokenText()

                token = parser.Scan()
                if parser.TokenText() == ":" {
                    
                    token = parser.Scan()
                    password := parser.TokenText()

                    // fmt.Println(min, max, char, password)

                    var entry password_entry
                    entry.min, _ = strconv.Atoi(min)
                    entry.max, _ = strconv.Atoi(max)
                    entry.char = rune(char[0])
                    entry.password = password

                    if is_entry_valid(entry) {
                        valid_password_counter_part_01 += 1
                    }

                    if is_entry_valid_part_02(entry) {
                        valid_password_counter_part_02 += 1
                    }
                }
            }
        } else {
            parsing = false
        }
    }

    fmt.Println("Amount of correct passwords part 01: ", valid_password_counter_part_01)
    fmt.Println("Amount of correct passwords part 02: ", valid_password_counter_part_02)
}

