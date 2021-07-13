package q16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	delta := math.MaxInt64
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}
	for k := 0; k < l; k++ {
		t := target - nums[k]
		i, j := k+1, l-1
		for i < j {
			d := t - nums[i] - nums[j]
			if d == 0 {
				return target
			} else if d < 0 {
				j--
			} else {
				i++
			}
			if abs(d) < abs(delta) {
				delta = d
			}
		}
	}

	return target - delta
}
