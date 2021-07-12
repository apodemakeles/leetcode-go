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

package q11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  []int
	output int
}{
	{
		input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		output: 49,
	},
	{
		input:  []int{1, 1},
		output: 1,
	},
	{
		input:  []int{4, 3, 2, 1, 4},
		output: 16,
	},
	{
		input:  []int{1, 2, 1},
		output: 2,
	},
}

func TestMaxArea(t *testing.T) {
	for i := range cases {
		actual := maxArea(cases[i].input)
		assert.Equal(t, cases[i].output, actual)
	}
}
