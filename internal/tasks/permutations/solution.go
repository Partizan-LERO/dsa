package permutations

import (
	"fmt"
)

func solution(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	if len(arr) == 2 {
		return 1
	}

	return len(arr) - 1*+len(arr)
}

func Solution() {
	fmt.Println(assertError(solution([]int{1, 2, 3}), 6))
	fmt.Println(assertError(solution([]int{1, 2, 3, 4}), 16))
}

func assertError(got, expected int) error {
	if got != expected {
		return fmt.Errorf("Got %s, expected %s", got, expected)
	}

	return nil
}
