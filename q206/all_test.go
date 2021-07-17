package q206

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	head   []int
	output []int
}{
	{
		head:   []int{1, 2, 3, 4, 5},
		output: []int{5, 4, 3, 2, 1},
	},
	{
		head:   []int{1, 2},
		output: []int{2, 1},
	},
	{
		head:   []int{},
		output: []int{},
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

func slice(list *ListNode) (s []int) {
	s = []int{}
	for list != nil {
		s = append(s, list.Val)
		list = list.Next
	}

	return
}

func TestReverseList(t *testing.T) {
	for i := range cases {
		actual := slice(reverseList(buildList(cases[i].head)))
		assert.Equal(t, cases[i].output, actual)
	}
}
