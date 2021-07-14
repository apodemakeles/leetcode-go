package q1818

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	nums1  []int
	nums2  []int
	output int
}{
	{
		nums1:  []int{1, 7, 5},
		nums2:  []int{2, 3, 5},
		output: 3,
	},
	{
		nums1:  []int{2, 4, 6, 8, 10},
		nums2:  []int{2, 4, 6, 8, 10},
		output: 0,
	},
	{
		nums1:  []int{1, 10, 4, 4, 2, 7},
		nums2:  []int{9, 3, 5, 1, 7, 4},
		output: 20,
	},
}

func TestMinAbsoluteSumDiff(t *testing.T) {
	for i := range cases {
		actual := minAbsoluteSumDiff(cases[i].nums1, cases[i].nums2)
		assert.Equal(t, cases[i].output, actual)
	}
}
