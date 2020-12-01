package lsort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 4, 2, 3, 8}
	nums = BubbleSortSmallToLarge(nums)
	fmt.Println(nums)
	nums = []int{5, 4, 2, 3, 8}
	nums = BubbleSortLargeToSmall(nums)
	fmt.Println(nums)
}
