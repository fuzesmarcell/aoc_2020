package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
)

type navigation_instruction struct {
	kind byte
	move_amount int
}

type v2i struct {
	x int
	y int
}

func ReadFileAndSplitOnNewLine(file_name string) []string {
	data, _ := ioutil.ReadFile(file_name)
	split_data := strings.Split(string(data), "\r\n")

	return split_data
}

func RotatePoint90Left(p v2i) v2i {
	result := v2i{p.y * -1, p.x}
	return result
}

func RotatePoint90Right(p v2i) v2i {
	result := v2i{p.y, p.x * -1}
	return result
}

func ManhattanLength(pos v2i) int {
	return int(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
}

func main() {
	split_data := ReadFileAndSplitOnNewLine("../../inputs/day_12.txt")
	instructions := make([]navigation_instruction, len(split_data))

	for i, instruction := range split_data {
		instructions[i].kind = instruction[0]
		instructions[i].move_amount, _ = strconv.Atoi(instruction[1:len(instruction)])
	}

	//            N
	//       W			E
	//			  S	
	//

	pos := v2i{0, 0}
	dir := v2i{1, 0}
	for _, nav_instruction := range instructions {
		switch nav_instruction.kind {
		case 'F':
			pos.x += (dir.x * nav_instruction.move_amount)
			pos.y += (dir.y * nav_instruction.move_amount)
		case 'L':
			for i := 0; i < nav_instruction.move_amount / 90; i++ {
				dir = RotatePoint90Left(dir)
			}
		case 'R':
			for i := 0; i < nav_instruction.move_amount / 90; i++ {
				dir = RotatePoint90Right(dir)
			}
		case 'N':
			pos.y += nav_instruction.move_amount
		case 'S':
			pos.y -= nav_instruction.move_amount
		case 'W':
			pos.x -= nav_instruction.move_amount
		case 'E':
			pos.x += nav_instruction.move_amount
		}
	}
	
	fmt.Println("Part 01:", ManhattanLength(pos))

	ship_pos := v2i{0, 0}
	// NOTE(fuzes): The waypoint is relative to the ship
	waypoint_pos := v2i{10, 1}
	for _, nav_instruction := range instructions {
		// fmt.Println("-------------------------------------------")
		// fmt.Printf("%c, %d\n", nav_instruction.kind, nav_instruction.move_amount)
		// fmt.Println("Old:", "Ship:", ship_pos, "Waypoint:", waypoint_pos)
		switch nav_instruction.kind {
		case 'F':
			ship_pos.x += waypoint_pos.x * nav_instruction.move_amount
			ship_pos.y += waypoint_pos.y * nav_instruction.move_amount
		case 'L':
			for i := 0; i < nav_instruction.move_amount / 90; i++ {
				waypoint_pos = RotatePoint90Left(waypoint_pos)
			}
		case 'R':
			for i := 0; i < nav_instruction.move_amount / 90; i++ {
				waypoint_pos = RotatePoint90Right(waypoint_pos)
			}
		case 'N':
			waypoint_pos.y += nav_instruction.move_amount
		case 'S':
			waypoint_pos.y -= nav_instruction.move_amount
		case 'W':
			waypoint_pos.x -= nav_instruction.move_amount
		case 'E':
			waypoint_pos.x += nav_instruction.move_amount
		}

		// fmt.Println("New:", "Ship:", ship_pos, "Waypoint:", waypoint_pos)
	}

	fmt.Println("Part 02:", ManhattanLength(ship_pos))
}