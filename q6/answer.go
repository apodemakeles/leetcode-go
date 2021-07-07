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

package q6

func convert(s string, numRows int) string {
	pos := 0
	ss := make([]byte, len(s))
	cycle := numRows*2 - 2
	if cycle == 0 {
		cycle = 1
	}
	for line := 0; line < numRows; line++ {
		i := line
		edge := line <= 0 || line == numRows-1
		pad := cycle - i*2
		for i < len(s) {
			ss[pos] = s[i]
			pos++
			if edge {
				i += cycle
			} else {
				i += pad
				pad = cycle - pad
			}
		}
	}

	return string(ss)
}
