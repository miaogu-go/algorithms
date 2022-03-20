package RemoveElement

//removeElement 移除元素
//https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { //nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
