package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//在每个树行中找最大值
//https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
func main() {
	fmt.Println(largestValues(tree1()))
	fmt.Println(largestValues(tree2()))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	list := make([]*TreeNode, 0)
	result := make([]int, 0)
	list = append(list, root)
	for len(list) > 0 {
		//取出队列里的所有元素
		//获取这一层最大的节点
		tmp := make([]*TreeNode, 0)
		max := math.MinInt
		for _, node := range list {
			if node.Val > max {
				max = node.Val
			}
			//同时放入下一层节点
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		result = append(result, max)
		list = tmp
	}

	return result
}

func tree1() *TreeNode {
	root := TreeNode{Val: 1}
	l := TreeNode{Val: 3}
	r := TreeNode{Val: 2}
	l2 := TreeNode{Val: 5}
	r2 := TreeNode{Val: 3}
	rr2 := TreeNode{Val: 9}

	root.Left = &l
	root.Right = &r
	l.Left = &l2
	l.Right = &r2
	r.Right = &rr2

	return &root
}

func tree2() *TreeNode {
	root := TreeNode{Val: 0}
	l := TreeNode{Val: -1}
	root.Left = &l

	return &root
}
