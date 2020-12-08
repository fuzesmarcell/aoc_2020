package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type instruction struct {
	executed bool
	operation string
	value int
}

type program_result struct {
	exit_code string
	accumulator int
	instruction_pointer int
}

func ExecuteProgram(program []instruction) program_result {
	result := program_result{"", 0, 0}

	// NOTE(fuzes): Reset the program
	for i, _ := range program {
		pi := &program[i]
		pi.executed = false
	}

	instruction_pointer := 0
	accumulator := 0
	is_running := true

	for is_running {
		if instruction_pointer == len(program) {
			result.exit_code = "clean"
			is_running = false
		} else {
			inst := &program[instruction_pointer]

			if inst.executed {
				result.exit_code = "loop"
				is_running = false
			} else {
				inst.executed = true
				if inst.operation == "nop" {
					instruction_pointer += 1
				} else if inst.operation == "acc" {
					accumulator += inst.value
					instruction_pointer += 1
				} else if inst.operation == "jmp" {
					instruction_pointer += inst.value
				}
			}
		}
	}

	result.accumulator = accumulator
	result.instruction_pointer = instruction_pointer

	return result
}

func main() {
	data, err := ioutil.ReadFile("../inputs/day_08.txt")
	if err == nil {
		instructions_split := strings.Split(string(data), "\r\n")
		program := make([]instruction, len(instructions_split))

		for i, line := range instructions_split {
			line_split := strings.Split(line, " ")
			instruction_name := line_split[0]
			value, _ := strconv.Atoi(line_split[1])
			program[i] = instruction{false, instruction_name, value}
		}

		run_result := ExecuteProgram(program)

		fmt.Println("Part 01:", run_result.accumulator)

		// NOTE(fuzes): Dirty dirty brute force technique I guess... =)
		for i, _ := range program {
			p_instruction := &program[i]
			if p_instruction.operation == "jmp" {
				p_instruction.operation = "nop"

				run_result = ExecuteProgram(program)
				if run_result.exit_code == "clean" {
					fmt.Println("Part 02:", run_result.accumulator)
					break
				}

				p_instruction.operation = "jmp"
			} else if p_instruction.operation == "nop" {
				p_instruction.operation = "jmp"

				run_result = ExecuteProgram(program)
				if run_result.exit_code == "clean" {
					fmt.Println("Part 02:", run_result.accumulator)
					break
				}

				p_instruction.operation = "nop"
			}
		}
	}
}