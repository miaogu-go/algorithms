package ReverseString

//reverseString 反转字符串
//https://leetcode-cn.com/problems/reverse-string/
func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	start, end := 0, len(s)-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
