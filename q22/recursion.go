package q22

func generateParenthesisRecursive(n int) []string {
	r := make([]string, 0)
	genRecursive(r, nil, n, n)

	return r
}

func genRecursive(r []string, s []byte, left, right int) {
	if left == 0 {
		for i := 0; i < right; i++ {
			s = append(s, ')')
		}
		r = append(r, string(s))

		return
	}
	i := len(s)
	genRecursive(r, append(s, '('), left-1, right)
	if left > right {
		genRecursive(r, append(s[:i], ')'), left, right-1)
	}
}
