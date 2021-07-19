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

package q23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  [][]int
	output []int
}{
	//{
	//	input:  [][]int{{}, {}},
	//	output: []int{},
	//},
	//{
	//	input:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
	//	output: []int{1, 1, 2, 3, 4, 4, 5, 6},
	//},
	//{
	//	input:  [][]int{},
	//	output: []int{},
	//},
	//{
	//	input:  [][]int{{}},
	//	output: []int{},
	//},
	//{
	//	input:  [][]int{{1}, {2}, {3}},
	//	output: []int{1, 2, 3},
	//},
	//{
	//	input:  [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
	//	output: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
	//},
	{
		input:  [][]int{{}, {-1, 1, 1, 5}, {}, {6, 10}},
		output: []int{-1, 1, 1, 5, 6, 10},
	},
}

func buildListList(input [][]int) []*ListNode {
	var nodes []*ListNode
	for i := range input {
		nodes = append(nodes, buildList(input[i]))
	}

	return nodes
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

func TestMergeKLists(t *testing.T) {
	for i := range cases {
		actual := slice(mergeKLists(buildListList(cases[i].input)))
		assert.Equal(t, cases[i].output, actual)
	}
}

func TestMergeKListsSort(t *testing.T) {
	for i := range cases {
		actual := slice(mergeKListsSort(buildListList(cases[i].input)))
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
