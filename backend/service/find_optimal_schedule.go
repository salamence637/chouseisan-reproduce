package service

import (
	"chouseisan/repository"
	"fmt"
	"log"
)

func createPreferencesCount(availability map[uint](map[uint](uint)), eventTimeslots []repository.EventTimeslot) map[uint](map[uint]uint) {
	preferenceCount := make(map[uint](map[uint]uint))
	slice := []uint{1, 2, 3}
	for _, value := range slice {
		preferenceCount[value] = make(map[uint]uint)
	}
	//initialization
	for _, eventTimeslot := range eventTimeslots {
		preferenceCount[1][eventTimeslot.ID] = 0
		preferenceCount[1][eventTimeslot.ID] = 1
		preferenceCount[1][eventTimeslot.ID] = 2
	}
	for _, userTimeslots := range availability {
		for timeslot, pref := range userTimeslots {
			if pref == 1 || pref == 2 || pref == 3 {
				preferenceCount[pref][timeslot]++
			}
		}
	}
	return preferenceCount
}

func intersection(arr1, arr2 []uint) []uint {
	// arr1 -> maxList, arr2 -> minList
	// Map to store the elements of the first array
	elements := make(map[uint]bool)

	// Store the elements of the first array in the map
	for _, elem := range arr1 {
		elements[elem] = true
	}

	// Slice to hold the intersection
	var intersect []uint

	if len(arr1) == 0 && len(arr2) == 0 {
		return intersect
	} else if len(arr2) == 0 {
		return arr1
	} else if len(arr1) == 0 {
		return arr2
	}

	// Iterate over the second array
	for _, elem := range arr2 {
		if _, found := elements[elem]; found {
			// Check if the element is in the map (and hence in the first array)
			intersect = append(intersect, elem)

			// avoid duplication
			delete(elements, elem)
		}
	}

	return intersect
}

func FindOptimals(availability map[uint](map[uint](uint)), n_users int, eventTimeslots []repository.EventTimeslot) ([]uint, []uint, []uint) {
	preferenceCount := createPreferencesCount(availability, eventTimeslots)
	log.Println(preferenceCount)
	fmt.Println(preferenceCount)
	// find schedule with the most 3's
	var maxTimeslots []uint
	var maxCount uint
	innerMap, exists := preferenceCount[3]
	if exists {
		for timeslot, count := range innerMap {
			if count > maxCount {
				maxCount = count
				var emptyArray []uint
				maxTimeslots = emptyArray
				maxTimeslots = append(maxTimeslots, timeslot)
			} else if count == maxCount {
				maxTimeslots = append(maxTimeslots, timeslot)
			}
		}
	}
	// find schedule with the least 1's
	var minTimeslots []uint
	minCount := uint(n_users)
	innerMap2, exists2 := preferenceCount[1]
	if exists2 {
		for timeslot, count := range innerMap2 {
			log.Println("count:", count)
			if count < minCount {
				minCount = count
				var emptyArray []uint
				minTimeslots = emptyArray
				minTimeslots = append(minTimeslots, timeslot)
			} else if count == minCount {
				minTimeslots = append(minTimeslots, timeslot)
			}
		}
	}

	log.Println(minTimeslots)
	fmt.Println(minTimeslots)

	// find intersection of both
	optimal := intersection(maxTimeslots, minTimeslots)

	return maxTimeslots, minTimeslots, optimal
}
