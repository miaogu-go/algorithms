package main

import (
	"fmt"
)

//实现 strStr()
//https://leetcode.cn/problems/implement-strstr/
func main() {

	pos := StrStr2("aaaaa", "bba")
	fmt.Println(pos)

}

func StrStr2(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if len(needle)+i > len(haystack) {
			return -1
		}
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

func StrStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	i, j := 0, 0
	length := len(haystack) - len(needle) + 1
	for i = 0; i < length; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}
