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

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var pre *ListNode
	carry := false
	for l1 != nil || l2 != nil || carry {
		sum := 0
		if carry {
			sum = 1
			carry = false
		}
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			sum -= 10
			carry = true
		}
		if pre == nil {
			pre = &ListNode{Val: sum}
			result = pre
		} else {
			pre.Next = &ListNode{Val: sum}
			pre = pre.Next
		}
	}

	return result
}
