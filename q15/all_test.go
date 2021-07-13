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

package q15

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  []int
	output [][]int
}{
	{
		input:  []int{-2, 0, 1, 1, 2},
		output: [][]int{{-2, 0, 2}, {-2, 1, 1}},
	},
	{
		input:  []int{-1, 0, 1, 2, -1, -4},
		output: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		input:  []int{},
		output: [][]int{},
	},
	{
		input:  []int{0},
		output: [][]int{},
	},
	{
		input:  []int{0},
		output: [][]int{},
	},
}

func TestThreeSum(t *testing.T) {
	for i := range cases {
		actual := threeSum(cases[i].input)
		equals(t, cases[i].output, actual)
	}
}

func equals(t *testing.T, expected, actual [][]int) {
	if len(expected) != len(actual) {
		assert.Fail(t, "the length is different")
	}
	for i := range actual {
		sort.Ints(actual[i])
	}
	for i := range expected {
		sort.Ints(expected[i])
		found := false
		for i := range actual {
			if reflect.DeepEqual(expected[i], actual[i]) {
				found = true
			}
		}
		if !found {
			assert.Fail(t, fmt.Sprintf("%v is not in actual result", expected[i]))
		}
	}

}
