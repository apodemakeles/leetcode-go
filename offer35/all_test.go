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

package offer35

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	nums   []int
	target int
	output int
}{
	{
		nums:   []int{5, 7, 7, 8, 8, 10},
		target: 8,
		output: 2,
	},
	{
		nums:   []int{5, 7, 7, 8, 8, 10},
		target: 6,
		output: 0,
	},
	{
		nums:   []int{},
		target: 6,
		output: 0,
	},
}

func TestSearch(t *testing.T) {
	for i := range cases {
		actual := search(cases[i].nums, cases[i].target)
		assert.Equal(t, cases[i].output, actual)
	}
}
