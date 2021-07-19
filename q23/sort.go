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

import "sort"

func mergeKListsSort(lists []*ListNode) *ListNode {
	last := len(lists) - 1
	sort.Slice(lists, func(i, j int) bool {
		if lists[i] == nil {
			return false
		} else if lists[j] == nil {
			return true
		}
		return lists[i].Val >= lists[j].Val
	})
	dummy := &ListNode{}
	p := dummy
	for last >= 0 {
		if lists[last] == nil {
			last--
			continue
		}
		p.Next = lists[last]
		p = p.Next
		if lists[last].Next == nil {
			last--
			continue
		}
		lists[last] = lists[last].Next

		if last > 0 {
			i := last
			tmp := lists[last]
			for ; i >= 1 && lists[last].Val > lists[i-1].Val; i-- {
			}
			for j := last - 1; j >= i; j-- {
				lists[j+1] = lists[j]
			}
			lists[i] = tmp
		}
	}

	return dummy.Next
}
