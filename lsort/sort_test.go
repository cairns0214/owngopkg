package lsort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 4, 2, 3, 8}
	nums = BubbleSort(nums, 0)
	fmt.Println(nums)
	nums = []int{5, 4, 2, 3, 8}
	nums = BubbleSort(nums, 1)
	fmt.Println(nums)
}
