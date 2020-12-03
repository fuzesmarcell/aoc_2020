package main

import (
	"fmt"
	"io/ioutil"
)

type slope struct {
	x_advance int
	y_advance int
}

func get_sled_grid_area(grid []byte, columns int, x int, y int) byte {
	
	// TODO(fuzes): Check the y value for range!
	// NOTE(fuzes): In the x direction the map wraps around!
	x_index := x % columns
	abs_index := y * columns + x_index
	result := grid[abs_index]
	return result
}

func main() {

	data, err := ioutil.ReadFile("../inputs/day_03_part_01.txt")
	if err == nil {
		number_of_rows := 0
		number_of_colums := 0
		for i, char := range data {
			if char == '\r' {
				number_of_colums = i
				break
			}
		}

		// NOTE:(fuzes): In the length of the data we have to calculate for the extra new line character at
		// the end of each row.
		number_of_new_line_chars := 2
		number_of_rows = (len(data) / (number_of_colums + number_of_new_line_chars))
		fmt.Println("Rows:", number_of_rows, "Colums:", number_of_colums)

		sled_grid := make([]byte, number_of_rows * number_of_colums)

		for y := 0; y < number_of_rows; y++ {
			for x := 0; x < number_of_colums; x++ {
				sled_grid[y * number_of_colums + x] = data[(y * (number_of_colums + number_of_new_line_chars)) + x]
			}
		}

		x := 0
		y := 0

		x_advance := 3
		y_advance := 1

		number_of_trees := 0
		for true {
			x += x_advance
			y += y_advance

			if y > number_of_rows - 1 {
				break
			}

			area := get_sled_grid_area(sled_grid, number_of_colums, x, y)
			// fmt.Printf("%d, %d, %c\n", x, y, area)
			if area == '#' {
				number_of_trees += 1
			}
		}

		fmt.Println("Part 01, Number of trees: ", number_of_trees)

		//
		// Part 02
		//

		x = 0
		y = 0
		tree_product := 1

		var slopes = []slope {
			slope {
				x_advance: 1,
				y_advance: 1,
			},
			slope {
				x_advance: 3,
				y_advance: 1,
			},
			slope {
				x_advance: 5,
				y_advance: 1,
			},
			slope {
				x_advance: 7,
				y_advance: 1,
			},
			slope {
				x_advance: 1,
				y_advance: 2,
			},
		}

		for _, advance_slope := range slopes {
			
			x = 0
			y = 0
			number_of_trees = 0

			for true {
				x += advance_slope.x_advance
				y += advance_slope.y_advance
	
				if y > number_of_rows - 1 {
					break
				}
	
				area := get_sled_grid_area(sled_grid, number_of_colums, x, y)
				// fmt.Printf("%d, %d, %c\n", x, y, area)
				if area == '#' {
					number_of_trees += 1
				}
			}

			fmt.Println("Number of trees:", number_of_trees)
			tree_product *= number_of_trees
		}

		fmt.Println("Part 02, Tree product: ", tree_product)
	}
}