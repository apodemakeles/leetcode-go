package q14

import "math"

func longestCommonPrefix(strs []string) (s string) {
	l := math.MaxInt64
	for _, str := range strs {
		if len(str) < l {
			l = len(str)
		}
	}
	if l == math.MaxInt32 {
		return ""
	}
	for i := 0; i < l; i++ {
		var x byte = 0
		for _, str := range strs {
			if x == 0 {
				x = str[i]
			} else if str[i] != x {
				return
			}
		}
		s += string(x)
	}

	return
}
