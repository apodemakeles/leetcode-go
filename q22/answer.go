package q22

func generateParenthesisRecursive(n int) []string {
	return genRecursive(nil, n, n)
}

func genRecursive(s []byte, left, right int) []string {
	var r []string
	if left == 0 {
		for i := 0; i < right; i++ {
			s = append(s, ')')
		}
		r = append(r, string(s))
	} else {
		i := len(s)
		r = append(r, genRecursive(append(s, '('), left-1, right)...)
		if left < right {
			r = append(r, genRecursive(append(s[:i], ')'), left, right-1)...)
		}
	}

	return r
}
