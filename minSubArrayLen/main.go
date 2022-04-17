package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

//minSubArrayLen 长度最小的子数组
//https://leetcode-cn.com/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	start, end := 0, 0
	sum := 0
	ans := math.MaxInt32
	for _, v := range nums {
		sum += v
		for sum >= target {
			if ans > end-start+1 {
				ans = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}
