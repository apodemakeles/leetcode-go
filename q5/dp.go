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
