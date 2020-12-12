package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	data, _ := ioutil.ReadFile("../../inputs/day_10.txt")
	split_data := strings.Split(string(data), "\r\n")
	joltage_ratings := make([]int, len(split_data))

	for i, snumber := range split_data {
		joltage_ratings[i], _ = strconv.Atoi(snumber)
	}

	charging_outlet_jolts := 0
	one_jolt_difference_counter := 0
	three_jolt_difference_counter := 0
	can_adapt := true
	for can_adapt {
		min_distance := 4
		next_jolt := 0
		for _, jolt := range joltage_ratings {
			distance := jolt - charging_outlet_jolts
			if distance > 0 && distance <= 3 {
				if distance < min_distance {
					min_distance = distance
					next_jolt = jolt
				}
			}
		}

		// fmt.Println(charging_outlet_jolts, next_jolt, min_distance)

		if next_jolt == 0 {
			can_adapt = false
		}

		if min_distance == 1 {
			one_jolt_difference_counter += 1
		}

		if min_distance == 3 {
			three_jolt_difference_counter += 1
		}

		charging_outlet_jolts = next_jolt
	}

	// NOTE(fuzes): We have to "count" the last adapter which is ours so its +3 counter
	three_jolt_difference_counter += 1

	fmt.Println("Part 01:", one_jolt_difference_counter, three_jolt_difference_counter,
							one_jolt_difference_counter * three_jolt_difference_counter)
}