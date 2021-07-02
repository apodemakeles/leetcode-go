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

package q1

import "sort"

func twoSumSort(nums []int, target int) []int {
	n := make([][2]int, len(nums))
	for i := range nums {
		n[i] = [2]int{nums[i], i}
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i][0] <= n[j][0]
	})

	i, j := 0, len(n)-1
	for i < j {
		sum := n[i][0] + n[j][0]
		if sum == target {
			return []int{n[i][1], n[j][1]}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return []int{}
}
