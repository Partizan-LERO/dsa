package first_and_last_indices_of_an_element_in_a_sorted_array

import "fmt"

func Solution() {
	list := []int{1, 2, 3, 4, 5, 7, 7, 9, 9, 15, 17}
	fmt.Println(solution(list, 9))
}

func solution(numbers []int, target int) []int {
	first := 0
	last := len(numbers) - 1

	firstIndex, lastIndex := 0, 0
	firstFound, lastFound := false, false

	for firstIndex <= last {
		if firstFound && lastFound {
			return []int{firstIndex, lastIndex}
		}
		if numbers[first] == target && !firstFound {
			firstIndex = first
			firstFound = true
		}

		if numbers[last] == target && !lastFound {
			lastIndex = last
			lastFound = true
		}

		last--
		first++
	}

	if firstFound && lastFound {
		return []int{firstIndex, lastIndex}
	}

	return make([]int, 0)
}
