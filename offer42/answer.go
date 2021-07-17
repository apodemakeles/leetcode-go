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

package offer42

func maxSubArray2(nums []int) int {
	max, total := -100, -100
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			if total < 0 {
				total = 0
			}
			total += nums[i]
		} else {
			if total > 0 {
				if total > max {
					max = total
				}
				total += nums[i]
			} else if nums[i] > max {
				max = nums[i]
			}
		}
	}
	if total > max {
		max = total
	}

	return max
}

func maxSubArray(nums []int) int {
	max, total := -100, -100
	for i := 0; i < len(nums); {
		if nums[i] >= 0 {
			if total < 0 {
				total = 0
			}
			total += nums[i]
			i++
			continue
		}
		if total > max {
			max = total
		}
		end := i
		for ; end < len(nums) && nums[end] < 0; end++ {
			total += nums[end]
			if nums[end] > max {
				max = nums[end]
			}
		}
		i = end
	}
	if total > max {
		max = total
	}

	return max
}
