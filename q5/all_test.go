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

package q5

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  string
	output []string
}{
	{
		input:  "a",
		output: []string{"a"},
	},
	{
		input:  "ac",
		output: []string{"a", "c"},
	},
	{
		input:  "cbbd",
		output: []string{"bb"},
	},
	{
		input:  "babad",
		output: []string{"bab", "aba"},
	},
}

func TestDP(t *testing.T) {
	for i := range cases {
		actual := longestPalindromeDP(cases[i].input)
		found := false
		for _, expected := range cases[i].output {
			if expected == actual {
				found = true
				break
			}
		}
		if !found {
			assert.Fail(t, fmt.Sprintf("expected %s, actual %s", cases[i].output, actual))
		}
	}
}
