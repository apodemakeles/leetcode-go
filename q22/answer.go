package q22

func generateParenthesis(n int) []string {
	var r []string
	var genRecursive func(s []byte, left, right int)
	genRecursive = func(s []byte, left, right int) {
		if left == 0 {
			for i := 0; i < right; i++ {
				s = append(s, ')')
			}
			r = append(r, string(s))
		} else {
			i := len(s)
			genRecursive(append(s, '('), left-1, right)
			if left < right {
				genRecursive(append(s[:i], ')'), left, right-1)
			}
		}
	}
	genRecursive(nil, n, n)

	return r
}
