package sorting

func BubbleSort(arr []int) []int {
	arrLength := len(arr)
	swapped := false

	for i := 0; i < arrLength; i++ {
		for j := 0; j < arrLength-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if swapped == false {
			return arr
		}
	}

	return arr
}
