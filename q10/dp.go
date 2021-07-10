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

package q10

func isMatchDP(s string, p string) bool {
	var reg []item
	for i := 0; i < len(p); i++ {
		c := p[i]
		if c == '*' {
			reg[len(reg)-1].w = true // panic if empty
		} else {
			reg = append(reg, item{c: c})
		}
	}
	if len(s) == 0 && len(reg) == 0 {
		return true
	} else if len(reg) == 0 {
		return false
	} else if len(s) == 0 {
		for _, r := range reg {
			if !r.w {
				return false
			}
		}
		return true
	}

	m := make([][]bool, len(reg))
	for i := range m {
		m[i] = make([]bool, len(s))
	}

	m[0][0] = match(s[0], reg[0])
	for i := 1; i < len(s) && m[0][i-1]; i++ {
		m[0][i] = reg[0].w && match(s[i], reg[0])
	}
	var w bool
	w = reg[0].w
	for j := 1; j < len(reg); j++ {
		m[j][0] = (m[j-1][0] && reg[j].w) || (w && match(s[0], reg[j]))
		w = w && reg[j].w
	}

	for j := 1; j < len(reg); j++ {
		for i := 1; i < len(s); i++ {
			if reg[j].w && (m[j-1][i] || (m[j][i-1] && match(s[i], reg[j]))) {
				m[j][i] = true
			} else if m[j-1][i-1] && match(s[i], reg[j]) {
				m[j][i] = true
			}
		}
	}

	return m[len(reg)-1][len(s)-1]
}
