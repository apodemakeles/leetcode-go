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

package q2

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
		l1:     []int{2, 4, 3},
		l2:     []int{5, 6, 4},
		output: []int{7, 0, 8},
	},
	{
		l1:     []int{0},
		l2:     []int{0},
		output: []int{0},
	},
	{
		l1:     []int{9, 9, 9, 9, 9, 9, 9},
		l2:     []int{9, 9, 9, 9},
		output: []int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}

func slice2List(s []int) *ListNode {
	var head, pre *ListNode
	for i := range s {
		node := &ListNode{Val: s[i]}
		if pre == nil {
			pre = node
			head = pre
		} else {
			pre.Next = node
			pre = pre.Next
		}
	}

	return head
}

func list2Slice(node *ListNode) (s []int) {
	for node != nil {
		s = append(s, node.Val)
		node = node.Next
	}

	return
}

func TestAddTwoNumbers(t *testing.T) {
	for i := range cases {
		l1 := slice2List(cases[i].l1)
		l2 := slice2List(cases[i].l2)
		list := addTwoNumbers(l1, l2)
		actual := list2Slice(list)
		assert.Equal(t, cases[i].output, actual)
	}
}
