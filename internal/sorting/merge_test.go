package sorting

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	a1 := []int{10, 13, 14}
	a2 := []int{11, 12, 13}
	res := merge(a1, a2)

	if len(res) != 6 {
		t.Error("Error!")
	}

	fmt.Println(res)
}
