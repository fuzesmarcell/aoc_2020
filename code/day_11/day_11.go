package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type v2i struct {
	x int
	y int
}

const (
	seat_empty byte = 'L'
	seat_occupied byte = '#'
	seat_floor byte = '.'
)

func read_file_and_split_on_lines(file_name string) []string {
	data, _ := ioutil.ReadFile(file_name)
	split_data := strings.Split(string(data), "\r\n")

	return split_data
}

func GetNumberOfAdjecentOccupiedSeats(seats []byte, rows int, columns int, x int, y int) int {

	result := 0
	
	var adjecent_vectors = []v2i {
		v2i{0, 1},
		v2i{1, 1},
		v2i{1, 0},
		v2i{1, -1},
		v2i{0, -1},
		v2i{-1, -1},
		v2i{-1, 0},
		v2i{-1, 1},
	}

	for _, offset := range adjecent_vectors {
		px := x + offset.x
		py := y + offset.y

		if px >= 0 &&
		   px < columns &&
		   py >= 0 &&
		   py < rows {
			   seat := seats[py * columns + px]
			   if seat == seat_occupied {
				   result += 1
			   }
		   }

	}

	return result
}

func IsInBounds(rows int, columns int, x int, y int) bool {
	if x >= 0 &&
	   x < columns &&
	   y >= 0 &&
	   y < rows {
		   return true
	}

	return false
}

func GetNumberOfAdjecentOccupiedSeatsRay(seats []byte, rows int, columns int, x int, y int) int {
	result := 0
	
	var adjecent_vectors = []v2i {
		v2i{0, 1},
		v2i{1, 1},
		v2i{1, 0},
		v2i{1, -1},
		v2i{0, -1},
		v2i{-1, -1},
		v2i{-1, 0},
		v2i{-1, 1},
	}

	for _, offset := range adjecent_vectors {
		px := x + offset.x
		py := y + offset.y
		for IsInBounds(rows, columns, px, py) {
			seat := seats[py * columns + px]

			if seat == seat_occupied {
				result += 1
				break
			} else if seat == seat_empty {
				// NOTE(fuzes): A empty seat means we can not look past it!!
				break
			}

			px += offset.x
			py += offset.y
		}
	}

	return result
}

func Part01() {
	split_data := read_file_and_split_on_lines("../../inputs/day_11.txt")

	rows := len(split_data)
	columns := len(split_data[0])
	fmt.Println(rows, columns)
	var seat_set [2][]byte
	seat_set[0] = make([]byte, rows * columns)
	seat_set[1] = make([]byte, rows * columns)

	for y, line := range split_data {
		for x, char := range line {
			seat_set[0][y * columns + x] = byte(char)
			seat_set[1][y * columns + x] = byte(char)
		}
	}

	working_seat_index := 0
	storage_seat_index := 1

	running := true
	loop_counts := 0
	for running {
		loop_counts += 1
		seat_changed := false
		for y := 0; y < rows; y++ {
			for x := 0; x < columns; x++ {
				seat := seat_set[working_seat_index][y * columns + x]

				switch seat {
				case seat_empty:
					number_of_adjecent_seats := GetNumberOfAdjecentOccupiedSeats(seat_set[working_seat_index],
																				 rows, columns, x, y)
					if number_of_adjecent_seats == 0 {
						seat_set[storage_seat_index][y * columns + x] = seat_occupied
						seat_changed = true
					}
				case seat_occupied:
					number_of_adjecent_seats := GetNumberOfAdjecentOccupiedSeats(seat_set[working_seat_index],
																				 rows, columns, x, y)
					if number_of_adjecent_seats >= 4 {
						seat_set[storage_seat_index][y * columns + x] = seat_empty
						seat_changed = true
					}
				}
			}
		}
		
		count := 0
		for y := 0; y < rows; y++ {
			for x := 0; x < columns; x++ {
				seat := seat_set[storage_seat_index][y * columns + x]
				if seat == seat_occupied {
					count += 1
				}
				// fmt.Printf("%c", seat)
			}

			// fmt.Printf("\n")
		}
		
		if working_seat_index == 0 {
			working_seat_index = 1
		} else {
			working_seat_index = 0
		}

		if storage_seat_index == 0 {
			storage_seat_index = 1
		} else {
			storage_seat_index = 0
		}

		for y := 0; y < rows; y++ {
			for x := 0; x < columns; x++ {
				seat_set[storage_seat_index][y * columns + x] = seat_set[working_seat_index][y * columns + x]
			}
		}


		running = seat_changed
		if !running {
			fmt.Println("Part 01:", count, loop_counts)
		}
	}
}

func Part02() {
	split_data := read_file_and_split_on_lines("../../inputs/day_11.txt")

	rows := len(split_data)
	columns := len(split_data[0])
	fmt.Println(rows, columns)
	var seat_set [2][]byte
	seat_set[0] = make([]byte, rows * columns)
	seat_set[1] = make([]byte, rows * columns)

	for y, line := range split_data {
		for x, char := range line {
			seat_set[0][y * columns + x] = byte(char)
			seat_set[1][y * columns + x] = byte(char)
		}
	}

	working_seat_index := 0
	storage_seat_index := 1

	running := true
	loop_counts := 0
	for running {
		loop_counts += 1
		seat_changed := false
		for y := 0; y < rows; y++ {
			for x := 0; x < columns; x++ {
				seat := seat_set[working_seat_index][y * columns + x]

				switch seat {
				case seat_empty:
					number_of_adjecent_seats := GetNumberOfAdjecentOccupiedSeatsRay(seat_set[working_seat_index],
																				    rows, columns, x, y)
					if number_of_adjecent_seats == 0 {
						seat_set[storage_seat_index][y * columns + x] = seat_occupied
						seat_changed = true
					}
				case seat_occupied:
					number_of_adjecent_seats := GetNumberOfAdjecentOccupiedSeatsRay(seat_set[working_seat_index],
																				    rows, columns, x, y)
					if number_of_adjecent_seats >= 5 {
						seat_set[storage_seat_index][y * columns + x] = seat_empty
						seat_changed = true
					}
				}
			}
		}
		
		count := 0
		for y := 0; y < rows; y++ {
			for x := 0; x < columns; x++ {
				seat := seat_set[storage_seat_index][y * columns + x]
				if seat == seat_occupied {
					count += 1
				}
				// fmt.Printf("%c", seat)
			}

			// fmt.Printf("\n")
		}

		// fmt.Printf("\n")
		
		if working_seat_index == 0 {
			working_seat_index = 1
		} else {
			working_seat_index = 0
		}

		if storage_seat_index == 0 {
			storage_seat_index = 1
		} else {
			storage_seat_index = 0
		}

		for y := 0; y < rows; y++ {
			for x := 0; x < columns; x++ {
				seat_set[storage_seat_index][y * columns + x] = seat_set[working_seat_index][y * columns + x]
			}
		}


		running = seat_changed
		if !running {
			fmt.Println("Part 02:", count, loop_counts)
		}
	}
}

func main() {
	// Part01()
	Part02()
}