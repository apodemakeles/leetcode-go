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

package q17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  string
	output []string
}{
	{
		input:  "23",
		output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
	{
		input:  "",
		output: []string{},
	},
	{
		input:  "2",
		output: []string{"a", "b", "c"},
	},
	{
		input:  "1",
		output: []string{},
	},
	{
		input:  "12",
		output: []string{"a", "b", "c"},
	},
}

func TestLetterCombinations(t *testing.T) {
	for i := range cases {
		actual := make(map[string]struct{})
		for _, s := range letterCombinations(cases[i].input) {
			actual[s] = struct{}{}
		}
		expected := make(map[string]struct{})
		for _, s := range cases[i].output {
			expected[s] = struct{}{}
		}
		assert.Equal(t, expected, actual)
	}
}
