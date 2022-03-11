package ReverseWords

import "strings"

//reverseWords 翻转字符串里的单词
//https://leetcode-cn.com/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	if s == "" {
		return ""
	}

	s = strings.TrimSpace(s)
	strs := strings.Split(s, " ")
	start, end := 0, len(strs)-1
	for start < end {
		strs[start], strs[end] = strs[end], strs[start]
		start++
		end--
	}

	result := make([]string, 0)
	for _, str := range strs {
		if str != "" {
			result = append(result, str)
		}
	}

	s = strings.Join(result, " ")

	return s
}
