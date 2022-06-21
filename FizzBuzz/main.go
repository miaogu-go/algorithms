package main

import (
	"fmt"
	"strconv"
)

/*给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：

answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
answer[i] == "Fizz" 如果 i 是 3 的倍数。
answer[i] == "Buzz" 如果 i 是 5 的倍数。
answer[i] == i （以字符串形式）如果上述条件全不满足。

输入：n = 3
输出：["1","2","Fizz"]

输入：n = 5
输出：["1","2","Fizz","4","Buzz"]

输入：n = 15
输出：["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]*/

func main() {
	s := fizzBuzz(3)
	fmt.Println(s)
	s = fizzBuzz(5)
	fmt.Println(s)
	s = fizzBuzz(15)
	fmt.Println(s)
}

func fizzBuzz(n int) []string {
	s := make([]string, 0)
	for i := 1; i <= n; i++ {
		tmp := ""
		if i%3 == 0 {
			tmp += "Fizz"
		}
		if i%5 == 0 {
			tmp += "Buzz"
		}
		if tmp == "" {
			tmp = strconv.Itoa(i)
		}
		s = append(s, tmp)
	}

	return s
}
