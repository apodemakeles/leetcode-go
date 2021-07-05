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

package q5

func longestPalindromeDP2(s string) string {
	l := len(s)
	m := make([][]bool, l)
	for i := range m {
		m[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		m[i][i] = true
	}

	for j := 0; j < l; j++ {
		for i := 0; i+j < l; i++ {
			if s[i] == s[i+j] && (j <= 2 || m[i+1][i+j-1]) {
				m[i][i+j] = true
			}
		}
	}

	for j := l - 1; j >= 0; j-- {
		for i := 0; i+j < l; i++ {
			if m[i][i+j] {
				return s[i : i+j+1]
			}
		}
	}

	return ""
}

func longestPalindromeDP(s string) string {
	lenMask := 0x03FF
	l := len(s)
	m := make([][]int, l)
	for i := range m {
		m[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		m[i][i] = i<<10 | 1
	}

	for j := 1; j < l; j++ {
		for i := 0; i+j < l; i++ {
			if s[i] == s[i+j] {
				p := false
				if i+1 < i+j-1 {
					val := m[i+1][i+j-1]
					if val>>10 == i+1 && (val&lenMask) == j-1 {
						p = true
					}
				} else {
					p = true
				}
				if p {
					m[i][i+j] = i<<10 | j + 1 // new full
					continue
				}
			}

			if m[i+1][i+j]&lenMask >= m[i][i+j-1]&lenMask {
				m[i][i+j] = m[i+1][i+j]
			} else {
				m[i][i+j] = m[i][i+j-1]
			}

		}
	}

	start := m[0][l-1] >> 10
	len := m[0][l-1] & lenMask

	return s[start : start+len]
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
