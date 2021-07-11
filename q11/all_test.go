package q11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  []int
	output int
}{
	{
		input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		output: 49,
	},
}

func TestIsMatchNormal(t *testing.T) {
	for i := range cases {
		actual := maxArea(cases[i].input)
		assert.Equal(t, cases[i].output, actual)
	}
}
