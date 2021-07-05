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

package q7

func reverse(x int) int {
	negative := x < 0
	var stack [10]int
	i := 0
	for {
		stack[i] = x % 10
		i++
		x /= 10
		if x == 0 {
			break
		}
	}

	var s []int
	if negative {
		s = []int{-2, -1, -4, -7, -4, -8, -3, -6, -4, -8}

	} else {
		s = []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
	}
	sum := 0
	o := i == len(s)
	for j := 0; j < i; j++ {
		if o { // overflow check
			a, b := stack[j], s[j]
			if negative {
				b, a = a, b
			}
			if a > b {
				return 0
			} else if a < b {
				o = false
			}
		}
		sum *= 10
		sum += stack[j]
	}

	return sum
}
