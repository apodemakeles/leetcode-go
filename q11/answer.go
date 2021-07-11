package q11

import "sort"

func maxArea(height []int) int {
	s := make([]int, len(height))
	for i := range s {
		s[i] = i
	}
	sort.Slice(s, func(i, j int) bool {
		return height[s[i]] > height[s[j]] || (height[s[i]] == height[s[j]] && s[i] > s[j])
	})
	left := []int{s[1]}
	right := []int{s[0]}
	l, r := s[1], s[0]
	for i := 2; i < len(s); i++ {
		if s[i] < l {
			l = s[i]
			left = append(left, s[i])
		} else if s[i] > r {
			r = s[i]
			right = append(right, s[i])
		}
	}

	min := func(h1, h2 int) int {
		if h1 <= h2 {
			return h1
		}

		return h2
	}
	area := func(i, j int) int {
		return (j - i) * min(height[i], height[j])
	}
	var max int
	for i := range left {
		for j := range right {
			if value := area(left[i], right[j]); value > max {
				max = value
			}
		}
	}

	return max
}
