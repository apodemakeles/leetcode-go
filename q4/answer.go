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

package q4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		return float64(findOne(nums1, nums2, l/2+1))
	}

	a, b := findTwo(nums1, nums2, l/2)
	return (float64(a) + float64(b)) / 2
}

func split(n1, n2 []int) (s1, s2, e1, e2 []int) {
	if len(n1) < len(n2) {
		n1, n2 = n2, n1
	}
	m1 := (len(n1) - 1) / 2

	i, j := 0, len(n2)-1
	for i < j {
		m2 := (i + j) / 2
		if n2[m2] >= n1[m1] {
			j = m2 - 1
		} else {
			i = m2 + 1
		}
	}
	if n1[m1] < n2[i] {
		i -= 1
	}

	s1 = n1[:m1+1]
	s2 = n2[:i+1]
	e1 = n1[m1+1:]
	e2 = n2[i+1:]

	return
}

func findTwo(n1, n2 []int, x int) (int, int) {
	for len(n1) > 0 && len(n2) > 0 {
		s1, s2, e1, e2 := split(n1, n2)
		left := len(s1) + len(s2)

		if left == x {
			return s1[len(s1)-1], findOne(e1, e2, x+1-left)
		} else if left == x+1 {
			return findOne(s1, s2, x), s1[len(s1)-1]
		} else if left > x+1 {
			n1, n2 = s1, s2
		} else {
			x -= left
			n1, n2 = e1, e2
		}
	}

	if len(n2) == 0 {
		return n1[x-1], n1[x]
	} else {
		return n2[x-1], n2[x]
	}
}

func min(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	}

	return n2
}

func findOne(n1, n2 []int, x int) int {
	for len(n1) > 0 && len(n2) > 0 {
		if len(n1) == 1 && len(n2) == 1 {
			return min(n1[0], n2[0])
		}

		s1, s2, e1, e2 := split(n1, n2)
		left := len(s1) + len(s2)

		if left == x {
			return s1[len(s1)-1]
		} else if left > x {
			n1, n2 = s1, s2
		} else {
			x -= left
			n1, n2 = e1, e2
		}
	}

	if len(n2) == 0 {
		return n1[x-1]
	} else {
		return n2[x-1]
	}
}
