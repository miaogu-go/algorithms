package LongestCommonPrefix

import (
	"unicode/utf8"
)

//longestCommonPrefix 最长公共前缀
//https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//获取最短字符串
	shortestLen := 99999
	shortestIndex := -1
	for i, str := range strs {
		strLen := utf8.RuneCountInString(str)
		if strLen < shortestLen {
			shortestLen = strLen
			shortestIndex = i
		}
	}
	if shortestIndex == -1 {
		return ""
	}

	//取最短值的每个元素进行比较，只要出现不匹配的直接返回已匹配的长度
	result := make([]byte, 0)
	shortestStr := strs[shortestIndex]
LOOP:
	for i := 0; i < utf8.RuneCountInString(shortestStr); i++ {
		for j, s := range strs {
			if j == shortestIndex {
				continue
			}
			if shortestStr[i] != s[i] {
				break LOOP
			}
		}
		result = append(result, shortestStr[i])
	}

	return string(result)
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	str0Len := len(strs[0])
	for i := 0; i < str0Len; i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
