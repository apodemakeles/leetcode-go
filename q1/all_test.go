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

package q1

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	nums   []int
	target int
	output []int
}{
	{
		nums:   []int{2, 7, 11, 15},
		target: 9,
		output: []int{0, 1},
	},
	{
		nums:   []int{3, 2, 4},
		target: 6,
		output: []int{1, 2},
	},
	{
		nums:   []int{3, 3},
		target: 6,
		output: []int{0, 1},
	},
}

func TestTwoSumSort(t *testing.T) {
	for i := range cases {
		actual := twoSumSort(cases[i].nums, cases[i].target)
		sort.Ints(actual)
		sort.Ints(cases[i].output)
		assert.Equal(t, cases[i].output, actual)
	}
}

func TestTwoSumMap(t *testing.T) {
	for i := range cases {
		actual := twoSumMap(cases[i].nums, cases[i].target)
		sort.Ints(actual)
		sort.Ints(cases[i].output)
		assert.Equal(t, cases[i].output, actual)
	}
}
