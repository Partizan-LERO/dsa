package sorting

func SelectionSort(arr []int) []int {
	arrLength := len(arr)
	for i := 0; i < arrLength; i++ {
		lowIndex := i
		for j := i + 1; j < arrLength; j++ {
			if arr[j] < arr[lowIndex] {
				lowIndex = j
			}
		}
		if lowIndex != i {
			arr[i], arr[lowIndex] = arr[lowIndex], arr[i]
		}
	}

	return arr
}
