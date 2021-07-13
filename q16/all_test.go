package q16

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
		nums:   []int{-1, 2, 1, -4},
		target: 1,
		output: 2,
	},
	{
		nums:   []int{5, 7, 0, -1, 1, 2, 11},
		target: 5,
		output: 5,
	},
	{
		nums:   []int{0, 0, 0},
		target: 5,
		output: 0,
	},
}

func TestThreeSumClosest(t *testing.T) {
	for i := range cases {
		actual := threeSumClosest(cases[i].nums, cases[i].target)
		assert.Equal(t, cases[i].output, actual)
	}
}
