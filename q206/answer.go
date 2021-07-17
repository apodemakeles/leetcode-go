package q206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	if head != nil {
		cur = head.Next
		head.Next = nil
	}
	for cur != nil {
		tmp := cur
		cur = cur.Next
		tmp.Next = head
		head = tmp
	}

	return head
}
