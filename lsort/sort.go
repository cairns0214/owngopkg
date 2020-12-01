// @Author: lvjinyu
// @Time: 2020/12/1 19:43
package lsort

// @Author lvjinyu
// @Description 冒泡排序从小到大排序
// @Date 20:11 2020/12/1
// @Param nums []int "需要排序的切片"
// @return nums []int "排序后的切片"
func BubbleSortSmallToLarge(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// @Author lvjinyu
// @Description 冒泡排序从大到小排序
// @Date 20:11 2020/12/1
// @Param nums []int "需要排序的切片"
// @return nums []int "排序后的切片"
func BubbleSortLargeToSmall(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
