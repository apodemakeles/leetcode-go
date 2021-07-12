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

package q11

import "sort"

func maxArea(height []int) int {
	s := make([]int, len(height))
	for i := range s {
		s[i] = i
	}
	sort.Slice(s, func(i, j int) bool {
		return height[s[i]] > height[s[j]] || (height[s[i]] == height[s[j]] && s[i] > s[j])
	})
	var l, r int
	if s[0] < s[1] {
		l, r = s[0], s[1]
	} else {
		r, l = s[0], s[1]
	}
	left := []int{l}
	right := []int{r}
	for i := 2; i < len(s); i++ {
		if s[i] < l {
			l = s[i]
			left = append(left, s[i])
		} else if s[i] >= r {
			r = s[i]
			right = append(right, s[i])
		}
	}

	min := func(h1, h2 int) int {
		if h1 <= h2 {
			return h1
		}

		return h2
	}
	area := func(i, j int) int {
		return (j - i) * min(height[i], height[j])
	}
	var max int
	for i := range left {
		for j := range right {
			if value := area(left[i], right[j]); value > max {
				max = value
			}
		}
	}

	return max
}
