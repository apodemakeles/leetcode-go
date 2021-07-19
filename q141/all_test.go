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

package q141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	head   []int
	pos    int
	output bool
}{
	{
		head:   []int{},
		pos:    -1,
		output: false,
	},
	{
		head:   []int{3, 2, 0, -4},
		pos:    1,
		output: true,
	},
	{
		head:   []int{1, 2},
		pos:    0,
		output: true,
	},
	{
		head:   []int{1},
		pos:    -1,
		output: false,
	},
	{
		head:   []int{1, 2},
		pos:    -1,
		output: false,
	},
}

func buildList(input []int, pos int) *ListNode {
	var head *ListNode
	var cur *ListNode
	var keep *ListNode
	for i := range input {
		node := &ListNode{Val: input[i]}
		if head == nil {
			head, cur = node, node
		} else {
			cur.Next = node
			cur = node
		}
		if i == pos {
			keep = node
		}
	}

	if keep != nil {
		cur.Next = keep
	}

	return head
}

func TestHasCycle(t *testing.T) {
	for i := range cases {
		actual := hasCycle(buildList(cases[i].head, cases[i].pos))
		assert.Equal(t, cases[i].output, actual)
	}
}
