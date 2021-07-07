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

package q9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	var s []uint8
	for x > 0 {
		s = append(s, uint8(x%10))
		x /= 10
	}
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isPalindromePlus(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	var s int64
	mask := int64(0x0F)
	l := 0
	for x > 0 {
		s = s<<4 + int64(x%10)
		l++
		x /= 10
	}
	i, j := 0, l-1
	for i <= j {
		if (s>>(4*i))&mask != (s>>(4*j))&mask {
			return false
		}
		i++
		j--
	}

	return true
}
