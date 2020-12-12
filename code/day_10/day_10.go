package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type FuncIntInt func([]int, int) int

func memorized(fn FuncIntInt) FuncIntInt {
	cache := make(map[int]int)

	return func(arr []int, input int) int {
		fmt.Println("Hello")
		if val, found := cache[input]; found {
			fmt.Println("Read from cache")
			return val
		}

		result := fn(arr, input)
		cache[input] = result
		return result
	}
}

func FindDeepArrangement(cache map[int]int, jolts []int, current_index int) int {

	// NOTE(fuzes): Memoization!!!
	if cached_result, is_in_hash := cache[current_index]; is_in_hash {
		// fmt.Println("Fetching from cache for index:", current_index)
		return cached_result
	}

	result := 0

	// NOTE(fuzes): In this case we hit the leaf node so we can count this one!
	if current_index == (len(jolts) - 1) {
		result = 1
	} else {
		for i := 1; i < 4; i++ {
			abs_i := current_index + i
			if abs_i < len(jolts) {
				if jolts[abs_i] - jolts[current_index] <= 3 {
					result += FindDeepArrangement(cache, jolts, abs_i)
				} else {
					break
				}
			} else {
				break
			}
		}
	}

	cache[current_index] = result

	return result
}

func main() {
	data, _ := ioutil.ReadFile("../../inputs/day_10.txt")
	split_data := strings.Split(string(data), "\r\n")
	joltage_ratings := make([]int, len(split_data))

	for i, snumber := range split_data {
		joltage_ratings[i], _ = strconv.Atoi(snumber)
	}

	var jolts_list []int
	jolts_list = append(jolts_list, 0)

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
		} else {
			jolts_list = append(jolts_list, next_jolt)
		}

		if min_distance == 1 {
			one_jolt_difference_counter += 1
		}

		if min_distance == 3 {
			three_jolt_difference_counter += 1
		}

		charging_outlet_jolts = next_jolt
	}

	jolts_list = append(jolts_list, jolts_list[len(jolts_list) - 1] + 3)

	// NOTE(fuzes): We have to "count" the last adapter which is ours so its +3 counter
	three_jolt_difference_counter += 1

	fmt.Println("Part 01:", one_jolt_difference_counter, three_jolt_difference_counter,
							one_jolt_difference_counter * three_jolt_difference_counter)


	cache := make(map[int]int)
	number_of_arrengemetns := FindDeepArrangement(cache, jolts_list, 0)
	fmt.Println("Part 02:", number_of_arrengemetns)
}