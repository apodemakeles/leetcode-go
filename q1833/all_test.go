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

package q1833

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	costs  []int
	coins  int
	output int
}{
	{
		costs:  []int{1, 3, 2, 4, 1},
		coins:  0,
		output: 0,
	},
	{
		costs:  []int{},
		coins:  10000,
		output: 0,
	},
	{
		costs:  []int{1, 3, 2, 4, 1},
		coins:  7,
		output: 4,
	},
	{
		costs:  []int{10, 6, 8, 7, 7, 8},
		coins:  5,
		output: 0,
	},
	{
		costs:  []int{1, 6, 3, 1, 2, 5},
		coins:  20,
		output: 6,
	},
	{
		costs:  []int{1, 9, 8, 4, 9, 7, 0, 9, 0, 7},
		coins:  100,
		output: 10,
	},
	{
		costs:  []int{27, 23, 33, 26, 46, 86, 70, 85, 89, 82, 57, 66, 42, 18, 18, 5, 46, 56, 42, 82, 52, 78, 4, 27, 96, 74, 74, 52, 2, 24, 78, 18, 42, 10, 12, 10, 80, 30, 60, 38, 32, 7, 98, 26, 18, 62, 50, 42, 15, 14, 32, 86, 93, 98, 47, 46, 58, 42, 74, 69, 51, 53, 58, 40, 66, 46, 65, 2, 10, 82, 94, 26, 6, 78, 2, 101, 97, 16, 12, 18, 71, 5, 46, 22, 58, 12, 18, 62, 61, 51, 2, 18, 34, 12, 36, 58, 20, 12, 17, 70},
		coins:  241,
		output: 24,
	},
}

func TestRecursive(t *testing.T) {
	for i := range cases {
		actual := maxIceCreamRecursive(cases[i].costs, cases[i].coins)
		assert.Equal(t, cases[i].output, actual)
	}
}

func TestSort(t *testing.T) {
	for i := range cases {
		actual := maxIceCreamSort(cases[i].costs, cases[i].coins)
		assert.Equal(t, cases[i].output, actual)
	}
}

func TestCounterSort(t *testing.T) {
	for i := range cases {
		actual := maxIceCreamCounterSort(cases[i].costs, cases[i].coins)
		assert.Equal(t, cases[i].output, actual)
	}
}
