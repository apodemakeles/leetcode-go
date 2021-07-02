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

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func maxIceCreamRecursive(costs []int, coins int) int {
	return maxIceCreamRecursiveInner(costs, 0, coins)
}

func maxIceCreamRecursiveInner(costs []int, i int, coins int) int {
	if coins == 0 {
		return 0
	}
	if i >= len(costs) {
		return 0
	}
	c := coins - costs[i]
	if c == 0 {
		return max(1, maxIceCreamRecursiveInner(costs, i+1, coins))
	} else if c < 0 {
		return maxIceCreamRecursiveInner(costs, i+1, coins)
	} else {
		return max(1+maxIceCreamRecursiveInner(costs, i+1, c),
			maxIceCreamRecursiveInner(costs, i+1, coins))
	}
}
