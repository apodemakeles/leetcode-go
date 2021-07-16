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

package q1846

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  []int
	output int
}{
	{
		input:  []int{2, 2, 1, 2, 1},
		output: 2,
	},
	{
		input:  []int{100, 1, 1000},
		output: 3,
	},
	{
		input:  []int{1, 2, 3, 4, 5},
		output: 5,
	},
	{
		input:  []int{73, 98, 9},
		output: 3,
	},
}

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	for i := range cases {
		actual := maximumElementAfterDecrementingAndRearranging(cases[i].input)
		assert.Equal(t, cases[i].output, actual)
	}
}
