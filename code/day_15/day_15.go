package main

import (
	"fmt"
)

type spoken_number_record struct {
	recent int
	last int
}

func WasNumberSpoken(records map[int]spoken_number_record, number int) bool {
	result := true

	if record, is_in_hash := records[number]; is_in_hash {
		result = record.recent > 0
	} else {
		result = false
	}

	return result
}

func main() {
	starting_numbers := []int {15,12,0,14,3,1}
	turn := 1

	record_map := make(map[int]spoken_number_record)

	for _, n := range starting_numbers{
		record_map[n] = spoken_number_record{0, turn}
		turn += 1
	}

	last_number_spoken := starting_numbers[len(starting_numbers) - 1]
	number_of_turns := 30000000
	for turn <= number_of_turns {
		if WasNumberSpoken(record_map, last_number_spoken) {
			record := record_map[last_number_spoken]
			last_number_spoken = record.last - record.recent
			
		} else {
			last_number_spoken = 0
		}

		// fmt.Println("Turn:", turn, "Spoken:", last_number_spoken)

		// NOTE(fuzes): Update the last and recent with the new number here
		rec := record_map[last_number_spoken]
		rec.recent = rec.last
		rec.last = turn
		record_map[last_number_spoken] = rec

		turn += 1
	}

	fmt.Println("Part 01:", last_number_spoken)

}