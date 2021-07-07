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

package q8

import "math"

func myAtoi(s string) (res int) {
	lead := true
	start := false
	positive := true
	for _, c := range s {
		if c >= '0' && c <= '9' {
			if lead {
				lead = false
			}
			if !start {
				start = true
			}

			num := c - '0'
			if positive {
				if (math.MaxInt32-num)/10 < int32(res) {
					return math.MaxInt32
				}
			} else {
				num = 0 - num
				if (math.MinInt32-num)/10 > int32(res) {
					return math.MinInt32
				}
			}
			res = res*10 + int(num)
		} else if c == ' ' && lead {
			continue
		} else if c == '+' && !start {
			lead, start, positive = false, true, true
		} else if c == '-' && !start {
			lead, start, positive = false, true, false
		} else {
			break
		}
	}

	return
}
