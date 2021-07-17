package q21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	l1     []int
	l2     []int
	output []int
}{
	{
		l1:     []int{1, 2, 4},
		l2:     []int{1, 3, 4},
		output: []int{1, 1, 2, 3, 4, 4},
	},
	{
		l1:     []int{},
		l2:     []int{},
		output: []int{},
	},
	{
		l1:     []int{},
		l2:     []int{0},
		output: []int{0},
	},
}

func buildList(input []int) *ListNode {
	var head *ListNode
	var cur *ListNode
	for i := range input {
		node := &ListNode{Val: input[i]}
		if head == nil {
			head, cur = node, node
		} else {
			cur.Next = node
			cur = node
		}
	}

	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	for i := range cases {
		actual := slice(mergeTwoLists(buildList(cases[i].l1), buildList(cases[i].l2)))
		assert.Equal(t, cases[i].output, actual)
	}
}

func slice(list *ListNode) (s []int) {
	s = []int{}
	for list != nil {
		s = append(s, list.Val)
		list = list.Next
	}

	return
}
