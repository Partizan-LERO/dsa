package sorting_with_three_unique_numbers

import "fmt"

func solution(numbers []int) []int {
	m := map[int]int{}
	m[1] = 0
	m[2] = 0
	m[3] = 0

	for _, n := range numbers {
		m[n] = m[n] + 1
	}

	var res []int
	x := 1
	for x <= 3 {
		i := 0
		for i < m[x] {
			res = append(res, x)
			i++
		}
		x++
	}

	return res
}

func Solution() {
	fmt.Println(solution([]int{1, 2, 3, 3, 3, 2, 1}))
}
