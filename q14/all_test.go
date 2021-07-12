package q14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  []string
	output string
}{
	{
		input:  []string{"flower", "flow", "flight"},
		output: "fl",
	},
	{
		input:  []string{"dog", "racecar", "car"},
		output: "",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for i := range cases {
		actual := longestCommonPrefix(cases[i].input)
		assert.Equal(t, cases[i].output, actual)
	}
}
