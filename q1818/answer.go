package q1818

import "sort"

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	l := len(nums1)
	s := make([]int, l, l)
	copy(s, nums1)
	sort.Ints(s)
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}

		return x
	}
	findMin := func(x int) (delta int) {
		i, j := 0, l-1
		for i < j {
			mid := (i + j) / 2
			y := s[mid]
			if x == y {
				return 0
			} else if x > y {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
		delta = abs(s[i] - x)
		if s[i] < x && i < l-1 {
			if d := abs(s[i+1] - x); d < delta {
				delta = d
			}
		} else if s[i] > x && i > 1 {
			if d := abs(s[i-1] - x); d < delta {
				delta = d
			}
		}

		return
	}
	var maxDel, sum int
	for i := range nums2 {
		delta := abs(nums2[i] - nums1[i])
		if delta == 0 {
			continue
		}
		sum += delta
		min := findMin(nums2[i])
		if delta-min > maxDel {
			maxDel = delta - min
		}
	}

	return (sum - maxDel) % 1000000007
}
