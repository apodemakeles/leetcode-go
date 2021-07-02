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

package q1833

import "sort"

func maxIceCreamSort(costs []int, coins int) int {
	sort.Slice(costs, func(i, j int) bool {
		if costs[i] <= costs[j] {
			return true
		}

		return false
	})

	total := 0
	for i := range costs {
		coins -= costs[i]
		if coins < 0 {
			break
		}
		total += 1
	}

	return total
}
