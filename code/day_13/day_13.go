package main

import (
	"fmt"
	"math"
)

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

	// var bus_times = []int64 {7,13,0,0,59,0,31,19}
	var bus_times = []int64 {17,0,13,19}
	// var bus_times = []int64 {67,7,59,61}
	// var bus_times = []int64 {67,0,7,59,61}
	// var bus_times = []int64 {67,7,0,59,61}
	// var bus_times = []int64 {1789,37,47,1889}
	// var bus_times = []int64 {29,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,41,0,0,0,0,0,0,0,0,0,521,0,0,0,0,0,0,0,23,0,0,0,0,13,0,0,0,17,0,0,0,0,0,0,0,0,0,0,0,0,0,601,0,0,0,0,0,37,0,0,0,0,0,0,0,0,0,0,0,0,19}
	// var bus_times = []int64 {29,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,41,0,0,0,0,0,0,0,0,0,521}


	k := int64(1)
	fmt.Println("Crunching away...")
	// k := int64(100000000000000 / bus_times[0])
	final_time_stamp := int64(0)
	sequence_found := false
	longest_sequence := 0
	for !sequence_found {

		next_time_stamp := bus_times[0] * k
		prev_i := 0
		for i := 1; i < len(bus_times); i++ {
			bus_id := bus_times[i]
			if bus_id != 0 {
				found, time_stamp := FindBusInTRange(next_time_stamp, bus_id, int64(i - prev_i))
				
				if !found {
					sequence_found = false
					break
				} else {
					if bus_id == 13 {
						fmt.Println(bus_times[0], k, bus_id, time_stamp / bus_id)
					}
					if i > longest_sequence {
						longest_sequence = i
						progress := float32(longest_sequence) / (float32(len(bus_times)) * 0.01)
						fmt.Printf("%f%% [%d, %d %d]\n", progress, longest_sequence, bus_id, k)
					}
				}

				prev_i = i
				next_time_stamp = time_stamp
			}

			sequence_found = true
		}

		final_time_stamp = bus_times[0] * k
		k += 1
	}

	fmt.Println("Part 02:", final_time_stamp, 100000000000000 - k)
}