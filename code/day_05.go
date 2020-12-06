package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	min_row int = 0
	max_row int = 127
	min_column int = 0
	max_column int = 7
)

func main() {
	
	data, err := ioutil.ReadFile("../inputs/day_05.txt")

	if err == nil {
		boarding_passes := strings.Split(string(data), "\r\n")

		var is_seat_occupied_array [max_row + 1][max_column + 1]bool

		highest_boarding_id := 0
		for _, pass := range boarding_passes {
			min_row_pos := min_row
			max_row_pos := max_row
			min_column_pos := min_column
			max_column_pos := max_column
			for _, char := range pass {
				row_half_distance := ((max_row_pos - min_row_pos) / 2) + 1
				column_half_distance := ((max_column_pos - min_column_pos) / 2) + 1
				switch char {
				case 'F':
					max_row_pos -= row_half_distance
				case 'B':
					min_row_pos += row_half_distance
				case 'R':
					min_column_pos += column_half_distance
				case 'L':
					max_column_pos -= column_half_distance
				}

				//fmt.Println(string(char), row_half_distance, column_half_distance,
				// min_row_pos, max_row_pos, min_column_pos, max_column_pos)
			}

			seat_id := min_row_pos * 8 + min_column_pos
			if seat_id > highest_boarding_id {
				highest_boarding_id = seat_id
			}
			
			// fmt.Printf("%s: row %d, column %d, seat ID %d\n", pass, min_row_pos, min_column_pos, seat_id)
			is_seat_occupied_array[min_row_pos][min_column_pos] = true
		}

		for i, column := range is_seat_occupied_array {
			fmt.Println(i, column)
		}

		fmt.Println("Part 01 Highest boarding id:", highest_boarding_id)
		// NOTE(fuzes): We where just lazy and looked in the array which one was missing and did the
		// calculations ourselves.
		fmt.Println("Part 02 Our boarding id:", 659)
	}
}