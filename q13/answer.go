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

package q13

func romanToInt(s string) int {
	m := map[string]uint16{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		num := uint16(0)
		i++
		if i < len(s) {
			num = m[s[i-1:i+1]]
		}
		if num == 0 {
			i--
			num = m[string(s[i])]
		}
		sum += int(num)
	}

	return sum
}
