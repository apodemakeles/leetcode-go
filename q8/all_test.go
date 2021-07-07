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

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  string
	output int
}{
	{
		input:  "42",
		output: 42,
	},
	{
		input:  "   -42",
		output: -42,
	},
	{
		input:  "4193 with words",
		output: 4193,
	},
	{
		input:  "words and 987",
		output: 0,
	},
	{
		input:  "-91283472332",
		output: math.MinInt32,
	},
	{
		input:  "3147483648",
		output: math.MaxInt32,
	},
}

func TestMyAtoi(t *testing.T) {
	for i := range cases {
		actual := myAtoi(cases[i].input)
		assert.Equal(t, cases[i].output, actual)
	}
}
