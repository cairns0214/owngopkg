// @Author: lvjinyu
// @Time: 2020/12/1 19:43
package lsort

// @Author lvjinyu
// @Description 冒泡排序
// @Date 20:11 2020/12/1
// @Param nums []int "需要排序的切片"
// @Param sortType int "排序类型, 0:从小到大, 1:从大到小"
// @return nums []int "排序后的切片"
func BubbleSort(nums []int, sortType int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	if sortType == 0 {
		for i := 0; i < length-1; i++ {
			for j := 0; j < length-i-1; j++ {
				if nums[j] > nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
	} else {
		for i := 0; i < length-1; i++ {
			for j := 0; j < length-i-1; j++ {
				if nums[j] < nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
	}
	return nums
}
