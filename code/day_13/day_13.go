package main

import (
	"fmt"
	"math"
)

type bus struct {
	id int64
	cof int64
}

type bus_pair struct {
	a bus
	b bus
}

func FindBusInTRange(bus_a_time_stamp int64, bus_b_id int64, t_range int64) (bool, int64) {
	
	bus_b_time_stamp := bus_b_id
	if bus_a_time_stamp > bus_b_time_stamp {
		cof := (bus_a_time_stamp / bus_b_time_stamp) + 1
		bus_b_time_stamp = bus_b_id * cof
	} else {
		return false, 0
	}

	distance := bus_b_time_stamp - bus_a_time_stamp
	
	if distance == t_range {
		return true, bus_b_time_stamp
	}

	return false, 0
}

func FindCofForBusPair(bus_a_id int64, bus_b_id int64, t_range int64) (int64, int64) {
	k := int64(1)
	for true {
		bus_a_time_stamp := bus_a_id * k
		found, time_stamp := FindBusInTRange(bus_a_time_stamp, bus_b_id, t_range)
		if found {
			return k, time_stamp / bus_b_id
		}

		k += 1
	}

	return 0, 0
}

func TestIfChainOfBusesIsValid(bus_times []int64, next_time_stamp int64) bool {

	sequence_found := false
	prev_i := 0
	for i := 1; i < len(bus_times); i++ {
		bus_id := bus_times[i]
		if bus_id != 0 {
			found, time_stamp := FindBusInTRange(next_time_stamp, bus_id, int64(i - prev_i))
			
			if !found {
				sequence_found = false
				break
			}

			prev_i = i
			next_time_stamp = time_stamp
		}

		sequence_found = true
	}

	return sequence_found
}

func GetFirstBusPairWithRange(bus_times []int64) (int64, int64, int64) {

	prev_i := 0
	for i := 1; i < len(bus_times); i++ {
		if bus_times[i] != 0 {
			return bus_times[0], bus_times[i], int64(i - prev_i)
		}
}

	return 0, 0, 0
}

func GetApproximationForEquation(a int64, a_inc int64, b int64, b_inc int64) (int64, int64, int64) {
	
	numbers := [2][3]int64 {{a, a_inc, 0}, {b, b_inc, 0}}
	
	max_iter := 1000000
	iter := 0
	for true {
		iter += 1

		min_index := 0
		for i := 1; i < len(numbers); i++ {
			if numbers[i - 1][0] < numbers[i][0] {
				min_index = i - 1
			} else if numbers[i - 1][0] > numbers[i][0] {
				min_index = i
			}
		}

		numbers[min_index][0] += numbers[min_index][1]
		numbers[min_index][2] += 1

		equal := false
		for i := 1; i < len(numbers); i++ {
			equal = numbers[i - 1][0] == numbers[i][0]
			if !equal {
				break
			}
		}

		if equal {
			// fmt.Println("Found with ither count:", iter)
			break
		}

		if iter > max_iter {
			fmt.Println("Max iteration exceeded!!!")
			return 0, 0, 0
		}
	}

	return numbers[0][0], numbers[0][2], numbers[1][2]
}

func GetTripletPairWithNewCof(pairs [4][2] int64) ([4][2]int64, [3]int64) {

	b_cof, x, y := GetApproximationForEquation(pairs[1][1], pairs[0][0], pairs[2][1], pairs[3][0])
	a_cof := pairs[1][0] * x + pairs[0][1]
	c_cof := pairs[2][0] * y + pairs[3][1]

	a := pairs[0][0] * a_cof
	b := pairs[1][0] * b_cof
	c := pairs[3][0] * c_cof

	// fmt.Println(a_cof, b_cof, c_cof)
	// fmt.Println(a, b, c)

	result_a := [4][2]int64 {{pairs[0][0], a_cof}, {pairs[1][0], b_cof}, {pairs[2][0], b_cof}, {pairs[3][0], c_cof}}
	result_b := [3]int64 {a, b, c}
	return result_a, result_b
}

func main() {
	our_departure_time := 1002461
	var bus_departure_times = []int {29, 41, 521, 23, 13, 17, 601, 37, 19}

	closest_depart_time := math.MaxInt64
	bus_id := 0
	for _, time_loop := range bus_departure_times {
		can_depart_with_bus := false
		t := 0
		for !can_depart_with_bus {
			t += time_loop
			if t >= our_departure_time {
				if t < closest_depart_time {
					closest_depart_time = t
					bus_id = time_loop
				}

				can_depart_with_bus = true
			}
		}
	}

	fmt.Println("Closest departure time:", closest_depart_time)
	part_01_solution := (closest_depart_time - our_departure_time) * bus_id
	fmt.Println("Part 01:", part_01_solution)
	
	pairs := [4][2] int64 {{7, 11}, {13, 6}, {13, 18}, {59, 4}}
	pairs_one := [4][2] int64 {pairs[2], pairs[3], {59, 11}, {31, 21}}
	a := int64(0)
	b := int64(0)
	for i := 0;; i++ {
		
		if i == 0 {
			var res_a [3]int64
			pairs, res_a = GetTripletPairWithNewCof(pairs)
			a = res_a[2]
	
			var res_b [3]int64
			pairs_one, res_b = GetTripletPairWithNewCof(pairs_one)
			b = res_b[1]
		}

		fmt.Println(a, b)

		if a == b {
			fmt.Println(pairs)
			fmt.Println(pairs_one)
			break
		}

		if a < b {
			var res_a [3]int64
			pairs, res_a = GetTripletPairWithNewCof(pairs)
			a = res_a[2]
		} else {
			var res_b [3]int64
			pairs_one, res_b = GetTripletPairWithNewCof(pairs_one)
			b = res_b[0]
		}
	}


	
	
	// // var bus_times = []int64 {7,13,0,0,59,0,31,19}
	// // var bus_times = []int64 {17,0,13,19}
	// // var bus_times = []int64 {67,7,59,61}
	// // var bus_times = []int64 {67,0,7,59,61}
	// // var bus_times = []int64 {67,7,0,59,61}
	// // var bus_times = []int64 {1789,37,47,1889}
	// var bus_times = []int64 {29,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,41,0,0,0,0,0,0,0,0,0,521,0,0,0,0,0,0,0,23,0,0,0,0,13,0,0,0,17,0,0,0,0,0,0,0,0,0,0,0,0,0,601,0,0,0,0,0,37,0,0,0,0,0,0,0,0,0,0,0,0,19}
	// // var bus_times = []int64 {29,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,41,0,0,0,0,0,0,0,0,0,521}


	// a_id, b_id, t_range := GetFirstBusPairWithRange(bus_times)
	// fmt.Println(a_id, b_id, t_range)
	// a_cof, b_cof := FindCofForBusPair(a_id, b_id, t_range)
	// fmt.Println(a_cof, b_cof)
	// // time_stamp := a_id * a_cof
	// time_stamp := int64(100000000000436)
	// for !TestIfChainOfBusesIsValid(bus_times, time_stamp) {
	// 	a_cof += b_id
	// 	time_stamp = a_id * a_cof
	// }

	// fmt.Println(time_stamp)
}