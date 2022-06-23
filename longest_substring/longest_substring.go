package longest_substring

/*func substringLen(s string) int {
	windowElem := make(map[byte]int, len(s))
	left, right, ans := 0, 0, 0
	for left < len(s) {
		if idx, ok := windowElem[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		windowElem[s[left]] = left
		left++
		ans = max(ans, left-right)
	}
	return ans
}*/

//https://leetcode.cn/problems/longest-substring-without-repeating-characters/
//substringLen 最长字符子串，滑动窗口解法
func substringLen(s string) int {
	left, right, leng := 0, 0, 0
	hash := make(map[uint8]int) //保存字符与字符位置的map，位置使用下标

	for right < len(s) {
		c, ok := hash[s[right]] //在map里查询是否存在字符
		if ok && c >= left {    //出现重复字符，重复字符的位置要大于左边的窗口
			left = c + 1 //移动左边下标到重复字符位置的后一位，让left、right之间没有重复字符
		}
		hash[s[right]] = right
		right++
		leng = max(leng, right-left)
	}

	return leng
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
