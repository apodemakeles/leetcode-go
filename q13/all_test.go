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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  string
	output int
}{
	{
		input:  "III",
		output: 3,
	},
	{
		input:  "IV",
		output: 4,
	},
	{
		input:  "IX",
		output: 9,
	},
	{
		input:  "LVIII",
		output: 58,
	},
	{
		input:  "MCMXCIV",
		output: 1994,
	},
}

func TestIntToRoman(t *testing.T) {
	for i := range cases {
		actual := romanToInt(cases[i].input)
		assert.Equal(t, cases[i].output, actual)
	}
}
