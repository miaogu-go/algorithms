package main

func main() {
	list := generateLink()
	deleteDuplicates(list)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateLink() *ListNode {
	arr := []int{1, 1, 2, 3, 3}
	var head, moveNode *ListNode

	for i, v := range arr {
		node := new(ListNode)
		node.Val = v
		if i == 0 {
			head = node
			moveNode = node
		}
		moveNode.Next = node
		moveNode = node
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	filter := make(map[int]struct{})
	node := head
	prevNode := head
	for node != nil {
		val := node.Val
		if _, ok := filter[val]; !ok {
			filter[val] = struct{}{}
			prevNode = node
		} else {
			prevNode.Next = node.Next
		}

		node = node.Next
	}
	return head
}
