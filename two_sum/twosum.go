package two_sum

/*给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]*/

func TwoSum(arr []int, target int) []int {
	tmp := make(map[int]int)
	for i, v := range arr {
		another := target - arr[i]
		index, ok := tmp[another]
		if ok {
			return []int{index, i}
		}
		tmp[v] = i
	}
	return nil
}

//twoSum 两数之和 II - 输入有序数组
//https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	startIndex := 0
	endIndex := len(numbers) - 1
	for startIndex < endIndex {
		sum := numbers[startIndex] + numbers[endIndex]
		if sum == target {
			return []int{startIndex + 1, endIndex + 1}
		}
		if sum < target {
			startIndex++
		} else {
			endIndex--
		}
	}

	return []int{-1, -1}
}
