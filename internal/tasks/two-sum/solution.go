package two_sum

import "fmt"

// you have a list of numbers and need to check if sum of any 2 numbers from it gives a target value.
// return indexes of these numbers or empty list
func Solution() {
	l := []int{2, 7, 11, 15}

	fmt.Println(solutionImproved(l, 18))
	fmt.Println(solutionImproved(l, 19))
	fmt.Println(solutionImproved(l, 17))
}

func solution(numbers []int, target int) []int {
	for i, num1 := range numbers {
		if num1 > target {
			continue
		}

		for j, num2 := range numbers {
			if num2 > target || i == j {
				continue
			}

			if num1+num2 == target {
				return []int{i, j}
			}
		}
	}

	return make([]int, 0)
}

// time improved
func solutionImproved(numbers []int, target int) []int {
	numbersMap := map[int]int{}
	for i, num := range numbers {
		numbersMap[num] = i
		if _, ok := numbersMap[target-num]; ok {
			return []int{numbersMap[target-num], numbersMap[num]}
		}
	}

	return make([]int, 0)
}
