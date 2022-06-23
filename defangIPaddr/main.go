package main

import (
	"fmt"
	"strings"
)

//IP 地址无效化
//https://leetcode.cn/problems/defanging-an-ip-address/

func main() {
	res := defangIPaddr("1.1.1.1")
	fmt.Println(res)
	res = defangIPaddr("255.100.50.0")
	fmt.Println(res)
}

func defangIPaddr(address string) string {
	ipArr := strings.Split(address, ".")
	result := strings.Join(ipArr, "[.]")
	return result
}
