package main

import "fmt"

//https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
//寻找旋转排序数组中的最小值
func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{1}))
	fmt.Println(findMin([]int{2, 1}))
	fmt.Println(findMin([]int{0, 1, 2}))
	fmt.Println(findMin([]int{2, 0, 1}))
	fmt.Println(findMin([]int{1, 2, 0}))
	fmt.Println(findMin([]int{0}))
}

func findMin(nums []int) int {
	if len(nums) < 2 && nums != nil {
		return nums[0]
	}

	first, second := 0, 1
	result := -5001
	for second <= len(nums)-1 {
		if nums[first] > nums[second] {
			result = nums[second]
			break
		}
		first++
		second++
	}

	if result == -5001 {
		return nums[0]
	}
	return result
}
