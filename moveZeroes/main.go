package main

import "fmt"

//https://leetcode.cn/problems/move-zeroes/
//移动零
func main() {
	num := []int{0, 1, 0, 3, 12}
	moveZeroes2(num)
	fmt.Println(num)

	num = []int{0}
	moveZeroes2(num)
	fmt.Println(num)

	num = []int{1, 0, 1}
	moveZeroes2(num)
	fmt.Println(num)
}

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	first, second := 0, 1
	for second <= len(nums)-1 {
		if nums[first] == 0 && nums[second] != 0 {
			nums[first], nums[second] = nums[second], nums[first]
			first++
		} else if nums[first] != 0 {
			first++
		}
		second++
	}
}

func moveZeroes2(nums []int) {
	if len(nums) < 2 {
		return
	}
	first, second := 0, 0
	for second <= len(nums)-1 {
		if nums[second] != 0 {
			nums[first], nums[second] = nums[second], nums[first]
			first++
		}
		second++
	}
}
