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

type item struct {
	w bool
	c byte
}

func equal(a, b byte) bool {
	return a == b || b == '.'
}

func exactlyMatch(s string, reg []item, i, j int) (bool, int, int) {
	for i < len(s) && j < len(reg) {
		pp := reg[j]
		if pp.w {
			break
		}
		if !equal(s[i], pp.c) {
			return false, -1, -1
		}
		i, j = i+1, j+1
	}

	return true, i, j
}

type stack struct {
	sis []int
	ris []int
}

func (s *stack) push(si, ri int) {
	s.sis = append(s.sis, si)
	s.ris = append(s.ris, ri)
}

func (s *stack) pop() (int, int, bool) {
	if len(s.sis) == 0 {
		return -1, -1, false
	}
	sis := s.sis[len(s.sis)-1]
	ris := s.ris[len(s.ris)-1]
	s.sis = s.sis[:len(s.sis)-1]
	s.ris = s.ris[:len(s.ris)-1]

	return sis, ris, true
}

func isMatchNormal(s string, p string) bool {
	var reg []item
	for i := 0; i < len(p); i++ {
		c := p[i]
		if c == '*' {
			reg[len(reg)-1].w = true // panic if empty
		} else {
			reg = append(reg, item{c: c})
		}
	}

	stack := &stack{}
	var i, j int
	for {
	L:
		for {
			if i >= len(s) {
				for jj := j; jj < len(reg); jj++ {
					if !reg[jj].w {
						break L
					}
				}
				return true
			}
			if j >= len(reg) && i < len(s) {
				break
			}

			pp := reg[j]
			if pp.w {
				if equal(s[i], pp.c) {
					stack.push(i, j)
				}
				j++
			} else {
				if match, ii, jj := exactlyMatch(s, reg, i, j); match {
					i, j = ii, jj
				} else { // mis-match
					break
				}
			}
		}

		var present bool
		if i, j, present = stack.pop(); present {
			i++
		} else {
			break
		}
	}

	return false
}
