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

package q10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	s      string
	p      string
	output bool
}{
	{
		s:      "aa",
		p:      "a",
		output: false,
	},
	{
		s:      "aa",
		p:      "a*",
		output: true,
	},
	{
		s:      "ab",
		p:      ".*",
		output: true,
	},
	{
		s:      "aab",
		p:      "c*a*b",
		output: true,
	},
	{
		s:      "mississippi",
		p:      "mis*is*p*.",
		output: false,
	},
	{
		s:      "mississippi",
		p:      "mis*is*ip*.",
		output: true,
	},
	{
		s:      "mississippi",
		p:      "mi.*i.*i.*.",
		output: true,
	},
	{
		s:      "ab",
		p:      ".*c",
		output: false,
	},
}

func TestIsMatchNormal(t *testing.T) {
	for i := range cases {
		actual := isMatchNormal(cases[i].s, cases[i].p)
		assert.Equal(t, cases[i].output, actual)
	}
}
