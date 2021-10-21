package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 2, 4, 1}
	fmt.Println(subSet(a))
}

func subSet2(nums []int) [][]int {
	ret := make([][]int, 0)
	list := make([]int, 0)
	backtrack(nums, 0, list, &ret)
	return ret
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {

}

func subSet(nums []int) [][]int {
	res := make([][]int, 0)
	revert := make(map[string]struct{})
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if _, ok := revert[fmt.Sprintf("%d-%d", nums[i], nums[j])]; ok {
				continue
			}
			data := make([]int, 0)
			if nums[i] == nums[j] {
				data = append(data, nums[i])
			} else {
				front, back := 0, 0
				if nums[i] < nums[j] {
					front = nums[i]
					back = nums[j]
				} else {
					front = nums[j]
					back = nums[i]
				}
				data = []int{front, back}
				revert[fmt.Sprintf("%d-%d", nums[j], nums[i])] = struct{}{}
			}
			res = append(res, data)
		}
	}
	if len(nums) > 2 {
		sort.Ints(nums)
		res = append(res, nums)
	}
	return res
}
