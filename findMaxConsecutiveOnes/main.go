package main

import "fmt"

func main() {
	one := []int{1, 1, 0, 1, 1, 1} //3
	//one := []int{1, 0, 1, 1, 0, 1} //2
	//one := []int{0} //0
	//one := []int{1, 1} //2
	//one := []int{0, 0} //0
	fmt.Println(findMaxConsecutiveOnes(one))
}

//findMaxConsecutiveOnes 最大连续1的个数
//https://leetcode-cn.com/problems/max-consecutive-ones/
func findMaxConsecutiveOnes(nums []int) int {
	oneLen := 0
	max := 0
	for _, v := range nums {
		if v == 1 {
			oneLen++
		} else {
			if oneLen > max {
				max = oneLen
			}
			oneLen = 0
		}
	}

	if oneLen > max {
		return oneLen
	}

	return max
}
