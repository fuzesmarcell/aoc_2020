package main

import (
	"fmt"
	"math"
	"pinhole"
)

type v3 struct {
	x int
	y int
	z int
}

func v3_add(a v3, b v3) v3 {
	return v3{a.x + b.x, a.y + b.y, a.z + b.z}
}

type cube struct {
	p         v3
	is_active bool
}

type space_cube_entry struct {
	cubes []cube
}

type v3_hash_table map[int]space_cube_entry

func SetSpaceCube(space v3_hash_table, p v3, active bool, allocate bool) {
	hash := p.x ^ p.y ^ p.z
	entry := space[hash]

	for i, c := range entry.cubes {
		if c.p.x == p.x && c.p.y == p.y && c.p.z == p.z {
			entry.cubes[i].is_active = active
			return
		}
	}

	if allocate {
		entry.cubes = append(entry.cubes, cube{p, active})
	}

	space[hash] = entry
}

func GetSpaceCube(space v3_hash_table, p v3) (bool, cube) {
	var result cube

	hash := p.x ^ p.y ^ p.z
	if _, is_in_hash := space[hash]; is_in_hash {
		entry := space[hash]

		for _, c := range entry.cubes {
			if c.p.x == p.x && c.p.y == p.y && c.p.z == p.z {
				return true, c
			}
		}
	}

	return false, result
}

func IsSpaceCubeActive(space v3_hash_table, p v3) bool {
	is_in_hash, cube := GetSpaceCube(space, p)
	if is_in_hash {
		return cube.is_active
	}

	return false
}

func GetNumberOfActiveCubesInNeighborhood(space v3_hash_table, p v3) int {
	offsets := []v3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},

		{1, 1, 0},
		{0, 1, 1},
		{1, 0, 1},

		{1, 1, 1},

		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},

		{-1, -1, 0},
		{0, -1, -1},
		{-1, 0, -1},

		{-1, -1, -1},

		{1, -1, 0},
		{-1, 1, 0},
		{0, 1, -1},
		{0, -1, 1},
		{-1, 0, 1},
		{1, 0, -1},

		{1, 1, -1},
		{-1, 1, 1},
		{1, -1, 1},

		{-1, -1, 1},
		{-1, 1, -1},
		{1, -1, -1},
	}

	result := 0
	for _, offset := range offsets {
		var pos v3
		pos.x = p.x + offset.x
		pos.y = p.y + offset.y
		pos.z = p.z + offset.z

		if IsSpaceCubeActive(space, pos) {
			result += 1
		}
	}

	return result
}

func CreateV3HashTableFromExisting(source v3_hash_table) v3_hash_table {
	dest := make(v3_hash_table)

	for key, value := range source {
		dest[key] = value
	}

	return dest
}

func PrintSpaceHashTable(space v3_hash_table, min v3, max v3) {
	for z := min.z; z <= max.z; z++ {
		fmt.Printf("z=%d\n", z)
		for y := min.y; y <= max.y; y++ {
			for x := min.x; x <= max.x; x++ {
				if IsSpaceCubeActive(space, v3{x, y, z}) {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}

			fmt.Printf("\n")
		}

		fmt.Printf("\n")
	}
}

func main() {

	space := make(v3_hash_table)

	initial_pos := ".#...####"
	row := 3
	column := 3

	min := v3{math.MaxInt32, math.MaxInt32, math.MaxInt32}
	max := v3{math.MinInt32, math.MinInt32, math.MinInt32}

	for y := 0; y < row; y++ {
		for x := 0; x < column; x++ {
			pos := v3{x, y, 0}
			switch initial_pos[y*column+x] {
			case '.':
				SetSpaceCube(space, pos, false, true)
			case '#':
				SetSpaceCube(space, pos, true, true)
			}

			if pos.x < min.x || pos.y < min.y || pos.z < min.z {
				min = pos
			}

			if pos.x > max.x || pos.y > max.y || pos.z > max.z {
				max = pos
			}
		}
	}

	min = v3_add(min, v3{-1, -1, -1})
	max = v3_add(max, v3{1, 1, 1})

	fmt.Println("Before any cycles:")
	PrintSpaceHashTable(space, min, max)

	space_storage := CreateV3HashTableFromExisting(space)

	number_of_cycles := 2
	for i := 0; i < number_of_cycles; i++ {
		next_min := min
		next_max := max
		for z := min.z; z <= max.z; z++ {
			for y := min.y; y <= max.y; y++ {
				for x := min.x; x <= max.x; x++ {
					pos := v3{x, y, z}
					number_of_active_cubes := GetNumberOfActiveCubesInNeighborhood(space, pos)
					if IsSpaceCubeActive(space, pos) {
						if number_of_active_cubes == 2 || number_of_active_cubes == 3 {
							SetSpaceCube(space_storage, pos, true, true)
						} else {
							SetSpaceCube(space_storage, pos, false, true)
						}

						// NOTE(fuzes): We have to expand our min max value so
						// that the search is greater for the next turn!
						if pos.x < next_min.x || pos.y < next_min.y || pos.z < next_min.z {
							next_min = pos
						}

						if pos.x > next_max.x || pos.y > next_max.y || pos.z > next_max.z {
							next_max = pos
						}

					} else {
						if number_of_active_cubes == 3 {
							SetSpaceCube(space_storage, pos, true, true)
						} else {
							SetSpaceCube(space_storage, pos, false, true)
						}
					}
				}
			}
		}

		fmt.Println(min, max)

		fmt.Println("Cycle:", i+1)
		PrintSpaceHashTable(space_storage, min, max)

		min = v3_add(next_min, v3{-1, -1, -1})
		max = v3_add(next_max, v3{1, 1, 1})

		space = CreateV3HashTableFromExisting(space_storage)
	}

	p := pinhole.New()
	p.DrawCube(-0.3, -0.3, -0.3, 0.3, 0.3, 0.3)
	p.DrawDot(0.2, 0.5, 0.5, 0.05)
	p.SavePNG("cube.png", 500, 500, nil)

	// number_of_cycles := 1
	// for i := 0; i < number_of_cycles; i++ {
	// 	for _, entry := range space {
	// 		for _, c := range entry.cubes {
	// 			number_of_active_cubes := GetNumberOfActiveCubesInNeighborhood(space, c.p)
	// 			fmt.Println(c, number_of_active_cubes)
	// 			if c.is_active {
	// 				if number_of_active_cubes == 2 || number_of_active_cubes == 3 {
	// 					SetSpaceCube(space_storage, c.p, true, true)
	// 				} else {
	// 					SetSpaceCube(space_storage, c.p, true, true)
	// 				}
	// 			} else {
	// 				if number_of_active_cubes == 3 {
	// 					SetSpaceCube(space_storage, c.p, true, true)
	// 				} else {
	// 					SetSpaceCube(space_storage, c.p, false, true)
	// 				}
	// 			}

	// 		}
	// 	}

	// 	space = CreateV3HashTableFromExisting(space_storage)
	// }

	// fmt.Println(GetNumberOfActiveCubesInNeighborhood(space, v3{1, 2, 0}))
}
