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

package q4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	num1   []int
	num2   []int
	output float64
}{
	{
		num1:   []int{1, 3},
		num2:   []int{2},
		output: 2,
	},
	{
		num1:   []int{1, 2},
		num2:   []int{3, 4},
		output: 2.5,
	},
	{
		num1:   []int{0, 0},
		num2:   []int{0, 0},
		output: 0,
	},
	{
		num1:   []int{},
		num2:   []int{1},
		output: 1,
	},
	{
		num1:   []int{2},
		num2:   []int{},
		output: 2,
	},
	{
		num1:   []int{3},
		num2:   []int{2},
		output: 2.5,
	},
	{
		num1:   []int{},
		num2:   []int{2, 3},
		output: 2.5,
	},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for i := range cases {
		actual := findMedianSortedArrays(cases[i].num1, cases[i].num2)
		assert.Equal(t, cases[i].output, actual)
	}
}
