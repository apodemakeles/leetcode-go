package q20

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	max := len(s) / 2
	stack := make([]byte, max, max)
	i := -1
	push := func(c byte) bool {
		i++
		if i >= max {
			return false
		}
		stack[i] = c

		return true
	}
	pop := func() (byte, bool) {
		if i < 0 {
			return 0, false
		}
		c := stack[i]
		i--

		return c, true
	}
	match := func(left, right byte) bool {
		return (left == '(' && right == ')') ||
			(left == '{' && right == '}') ||
			(left == '[' && right == ']')
	}

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			if !push(byte(c)) {
				return false
			}
		} else if b, ok := pop(); !ok || !match(b, byte(c)) {
			return false
		}
	}

	return i < 0
}
