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

func reverseList2(head *ListNode) *ListNode {
	cur := head
	var prv *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = prv
		prv = cur
		cur = tmp
	}

	return prv
}
