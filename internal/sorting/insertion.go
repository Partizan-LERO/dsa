package sorting

func InsertionSort(arr []int) []int {
	arrLength := len(arr)

	for i := 1; i < arrLength; i++ {
		current := arr[i]
		j := i

		for ; j >= 1 && arr[j-1] > current; j-- {
			arr[j] = arr[j-1]
		}

		arr[j] = current
	}

	return arr
}
