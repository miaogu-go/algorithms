package longest_substring

func substringLen(s string) int {
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
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
