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

package q22

import (
	"testing"

	"github.com/apodemakeles/leetcode-go/common"
)

var cases = []struct {
	n      int
	output []string
}{
	{
		n:      1,
		output: []string{"()"},
	},
	{
		n:      2,
		output: []string{"()()", "(())"},
	},
	{
		n:      3,
		output: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	},
	{
		n:      4,
		output: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
	},
}

func TestGenerateParenthesisRecursive(t *testing.T) {
	for i := range cases {
		actual := generateParenthesis(cases[i].n)
		common.StringSliceEquals(t, cases[i].output, actual)
	}
}
