package sorting

// MergeSort sorts an array in an ascending order
func MergeSort(arr []int) []int {
	pivot := len(arr) / 2

	if len(arr) < 2 {
		return arr
	}

	l := MergeSort(arr[0:pivot])
	r := MergeSort(arr[pivot:])

	return merge(l, r)
}

func merge(arr1, arr2 []int) []int {
	var res []int

	l1 := len(arr1)
	l2 := len(arr2)
	counter1 := 0
	counter2 := 0

	for i := 0; i < l1+l2; i++ {
		if counter1 == l1 && counter2 == l2 {
			break
		}

		if counter1 == l1 && counter2 != l2 {
			res = append(res, arr2[counter2:]...)
			break
		}

		if counter2 == l2 && counter1 != l1 {
			res = append(res, arr1[counter1:]...)
			break
		}

		if arr1[counter1] > arr2[counter2] {
			res = append(res, arr2[counter2])
			counter2++
			continue
		}

		if arr1[counter1] <= arr2[counter2] {
			res = append(res, arr1[counter1])
			counter1++
			continue
		}
	}

	return res
}
