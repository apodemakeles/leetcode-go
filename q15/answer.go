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

package q15

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	result := make([][]int, 0)
	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k-1] == nums[k] {
			continue
		}
		i, j := k+1, l-1
		for i < j {
			if i > k+1 && nums[i-1] == nums[i] {
				i++
				continue
			} else if j < len(nums)-1 && nums[j] == nums[j+1] {
				j--
				continue
			}
			v := nums[k] + nums[i] + nums[j]
			if v == 0 {
				result = append(result, []int{nums[k], nums[i], nums[j]})
				i++
			} else if v > 0 {
				j--
			} else {
				i++
			}
		}
	}

	return result
}
