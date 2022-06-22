package main

import "fmt"

//找树左下角的值
//https://leetcode.cn/problems/find-bottom-left-tree-value/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(findBottomLeftValue(tree1()))
	fmt.Println(findBottomLeftValue(tree2()))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	//临时队列
	list := make([]*TreeNode, 0)
	ans := 0
	list = append(list, root)
	for len(list) > 0 {
		//取出队首元素
		node := list[0]
		//更新队列
		list = list[1:]
		if node.Right != nil {
			list = append(list, node.Right)
		}
		if node.Left != nil {
			list = append(list, node.Left)
		}
		ans = node.Val
	}

	return ans
}

func tree1() *TreeNode {
	root := TreeNode{Val: 2}
	l1 := TreeNode{Val: 1}
	r1 := TreeNode{Val: 3}

	root.Left = &l1
	root.Right = &r1

	return &root
}

func tree2() *TreeNode {
	root := TreeNode{Val: 1}
	l1 := TreeNode{Val: 2}
	l2 := TreeNode{Val: 4}
	r1 := TreeNode{Val: 3}
	rl2 := TreeNode{Val: 5}
	rl3 := TreeNode{Val: 7}
	rr2 := TreeNode{Val: 6}

	rl2.Left = &rl3
	r1.Left = &rl2
	r1.Right = &rr2
	l1.Left = &l2
	root.Left = &l1
	root.Right = &r1

	return &root
}
