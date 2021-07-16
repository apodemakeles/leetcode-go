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

package q18

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	result := make([][]int, 0)
	for q := 0; q < l-3; q++ {
		if q > 0 && nums[q-1] == nums[q] {
			continue
		}
		for p := q + 1; p < l-2; p++ {
			if p > q+1 && nums[p-1] == nums[p] {
				continue
			}

			i, j := p+1, l-1
			for i < j {
				if i > p+1 && nums[i-1] == nums[i] {
					i++
					continue
				} else if j < len(nums)-1 && nums[j] == nums[j+1] {
					j--
					continue
				}
				v := nums[q] + nums[p] + nums[i] + nums[j]
				if v == target {
					result = append(result, []int{nums[q], nums[p], nums[i], nums[j]})
					i++
				} else if v > target {
					j--
				} else {
					i++
				}
			}
		}
	}

	return result
}

func fourSumPrune(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	result := make([][]int, 0)
	for q := 0; q < l-3; q++ {
		if q > 0 && nums[q-1] == nums[q] {
			continue
		}
		if nums[q]+nums[q+1]+nums[q+2]+nums[q+3] > target {
			break
		}
		if nums[q]+nums[l-1]+nums[l-2]+nums[l-3] < target {
			continue
		}
		for p := q + 1; p < l-2; p++ {
			if p > q+1 && nums[p-1] == nums[p] {
				continue
			}

			i, j := p+1, l-1
			for i < j {
				if i > p+1 && nums[i-1] == nums[i] {
					i++
					continue
				} else if j < len(nums)-1 && nums[j] == nums[j+1] {
					j--
					continue
				}
				v := nums[q] + nums[p] + nums[i] + nums[j]
				if v == target {
					result = append(result, []int{nums[q], nums[p], nums[i], nums[j]})
					i++
				} else if v > target {
					j--
				} else {
					i++
				}
			}
		}
	}

	return result
}
