package FindDiagonalOrder

//findDiagonalOrder 对角线遍历
//https://leetcode-cn.com/problems/diagonal-traverse/
func findDiagonalOrder(mat [][]int) []int {
	n, m := len(mat), len(mat[0])
	result := make([]int, 0, m*n)
	for i := 0; i < m+n-1; i++ {
		tmp := make([]int, 0)
		//确定对角线遍历位置
		//r 行，c 列
		r, c := 0, 0
		if i < m {
			r = 0
		} else {
			r = i - m + 1
		}
		if i < m {
			c = i
		} else {
			c = m - 1
		}

		//遍历对角线
		for r < m && c > -1 && r < n {
			tmp = append(tmp, mat[r][c])
			r++
			c--
		}

		//反转
		if i%2 == 0 {
			result = append(result, reservesArr(tmp)...)
		} else {
			result = append(result, tmp...)
		}

	}

	return result
}

//reservesArr 反转数组
func reservesArr(v []int) []int {
	if len(v) == 0 {
		return nil
	}
	startIndex, endIndex := 0, len(v)-1
	for startIndex < endIndex {
		v[startIndex], v[endIndex] = v[endIndex], v[startIndex]
		startIndex++
		endIndex--
	}

	return v
}
