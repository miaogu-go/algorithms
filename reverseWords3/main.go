package main

import (
	"fmt"
	"strings"
)

//反转字符串中的单词 III
//https://leetcode.cn/problems/reverse-words-in-a-string-iii/
func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("God Ding"))
}

func reverseWords(s string) string {
	if len(s) == 1 {
		return s
	}

	result := make([]string, 0)
	sArr := strings.Split(s, " ")
	for _, item := range sArr {
		reverStr := reverse(item)
		result = append(result, reverStr)
	}

	return strings.Join(result, " ")
}

func reverse(s string) string {
	if len(s) == 1 {
		return s
	}

	sBytes := []byte(s)
	first, last := 0, len(sBytes)-1
	for first <= last {
		sBytes[first], sBytes[last] = sBytes[last], sBytes[first]
		first++
		last--
	}

	return string(sBytes)
}
