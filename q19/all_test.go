/*
Copyright 2021 caozheng@troila.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package q19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  []int
	n      int
	output []int
}{
	{
		input:  []int{1, 2, 3, 4, 5},
		n:      2,
		output: []int{1, 2, 3, 5},
	},
	{
		input:  []int{1},
		n:      1,
		output: []int{},
	},
	{
		input:  []int{1, 2},
		n:      1,
		output: []int{1},
	},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for i := range cases {
		actual := slice(removeNthFromEnd(buildList(cases[i].input), cases[i].n))
		assert.Equal(t, cases[i].output, actual)
	}
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
