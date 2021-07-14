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

package q17

func letterCombinations(digits string) []string {
	m := map[int32][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "z"},
	}
	n := 1
	for _, x := range digits {
		if cs, present := m[x]; present {
			n *= len(cs)
		}
	}
	if n == 1 {
		return []string{}
	}
	r := make([]string, n, n)
	round := 1
	for i := range digits {
		num := digits[i]
		if cs, present := m[int32(num)]; present {
			t := n / (len(cs) * round)
			j := 0
			for m := round; m > 0; m-- {
				for _, c := range cs {
					for k := t; k > 0; k-- {
						r[j] = r[j] + c
						j++
					}
				}
			}
			round *= len(cs)
		}
	}

	return r
}
