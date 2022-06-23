package main

import "fmt"

//杨辉三角
//https://leetcode.cn/problems/pascals-triangle/

func main() {
	res := generate(5)
	fmt.Println(res)
	res = generate(1)
	fmt.Println(res)

	res = generate2(5)
	fmt.Println(res)
	res = generate2(1)
	fmt.Println(res)

	result := getRow(0)
	fmt.Println(result)
	result = getRow(1)
	fmt.Println(result)
	result = getRow(3)
	fmt.Println(result)
}

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 1; i <= numRows; i++ {
		if i <= 2 {
			tmp := make([]int, 0)
			for j := 0; j < i; j++ {
				tmp = append(tmp, 1)
			}
			result = append(result, tmp)
			continue
		}

		tmp := make([]int, i)
		tmp[0], tmp[i-1] = 1, 1
		//取出i-1的数组
		prev := result[i-2]

		x, y := 0, 1
		for z := 1; z <= i-2; z++ {
			tmp[z] = prev[x] + prev[y]
			x++
			y++
		}
		result = append(result, tmp)
	}

	return result
}

func generate2(numRows int) [][]int {
	results := make([][]int, numRows)
	for i := range results {
		results[i] = make([]int, i+1)
		results[i][0] = 1
		results[i][i] = 1
		for j := 1; j < i; j++ {
			results[i][j] = results[i-1][j] + results[i-1][j-1]
		}
	}

	return results
}

func getRow(rowIndex int) []int {
	result := make([][]int, 0)
	for i := 1; i <= rowIndex+1; i++ {
		if i <= 2 {
			tmp := make([]int, 0)
			for j := 0; j < i; j++ {
				tmp = append(tmp, 1)
			}
			result = append(result, tmp)
			continue
		}

		tmp := make([]int, i)
		tmp[0], tmp[i-1] = 1, 1
		//取出i-1的数组
		prev := result[i-2]

		x, y := 0, 1
		for z := 1; z <= i-2; z++ {
			tmp[z] = prev[x] + prev[y]
			x++
			y++
		}
		result = append(result, tmp)
	}

	return result[rowIndex]
}
