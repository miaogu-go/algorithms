package add_two_numbers

type Node struct {
	data int
	next *Node
}

//[3,4,5] + [4,5,6]
//7911
func addTwoNumber(l1, l2 *Node) *Node {
	carry := 0
	var currentNode *Node
	var head *Node
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.data
			l1 = l1.next
		}
		if l2 != nil {
			n2 = l2.data
			l2 = l2.next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &Node{data: sum}
			currentNode = head
		} else {
			currentNode.next = &Node{data: sum}
			currentNode = currentNode.next
		}
	}
	if carry > 0 {
		currentNode.next = &Node{data: carry}
	}
	return head
}
