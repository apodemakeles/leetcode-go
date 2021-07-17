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

package offer42

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	nums   []int
	output int
}{
	{
		nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		output: 6,
	},
	{
		nums:   []int{1},
		output: 1,
	},
	{
		nums:   []int{-1},
		output: -1,
	},
	{
		nums:   []int{-5, -4, -3, -2, -1},
		output: -1,
	},
	{
		nums:   []int{-5, -4, -3, 0, -2, -1},
		output: 0,
	},
	{
		nums:   []int{0},
		output: 0,
	},
	{
		nums:   []int{5, -4, -1, 0, 1, -1},
		output: 5,
	},
	{
		nums:   []int{5, -4, -2, 0, 6, -1},
		output: 6,
	},
	{
		nums:   []int{1, -2, -2, 0, -2, 2, -2, -2},
		output: 2,
	},
	{
		nums:   []int{1, 2, 3, 4, 5, 6},
		output: 21,
	},
}

func TestMaxSubArray(t *testing.T) {
	for i := range cases {
		actual := maxSubArray(cases[i].nums)
		assert.Equal(t, cases[i].output, actual)
	}
}

func TestMaxSubArray2(t *testing.T) {
	for i := range cases {
		actual := maxSubArray2(cases[i].nums)
		assert.Equal(t, cases[i].output, actual)
	}
}
