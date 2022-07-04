package main

import (
	"fmt"
	"math"
	"sort"
)

//最小绝对差
//https://leetcode.cn/problems/minimum-absolute-difference/
func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	ans := make([][]int, 0)
	best := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		delta := arr[i+1] - arr[i]
		if delta < best {
			best = delta
			ans = [][]int{{arr[i], arr[i+1]}}
		} else if delta == best {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}

	return ans
}
