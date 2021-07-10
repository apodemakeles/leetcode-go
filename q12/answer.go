package q12

func intToRomanMap(num int) (s string) {
	m := []struct {
		i uint16
		r string
	}{
		{i: 1000, r: "M"},
		{i: 900, r: "CM"},
		{i: 500, r: "D"},
		{i: 400, r: "CD"},
		{i: 100, r: "C"},
		{i: 90, r: "XC"},
		{i: 50, r: "L"},
		{i: 40, r: "XL"},
		{i: 10, r: "X"},
		{i: 9, r: "IX"},
		{i: 5, r: "V"},
		{i: 4, r: "IV"},
		{i: 1, r: "I"},
	}
	n := uint16(num)
	for n > 0 {
		for i := range m {
			if n >= m[i].i {
				n -= m[i].i
				s += m[i].r
				break
			}
		}
	}

	return
}

func intToRoman(num int) (s string) {
	if num >= 1000 {
		n := num / 1000
		for i := 0; i < n; i++ {
			s += "M"
		}
		num = num % 1000
	}
	if num >= 100 {
		n := num / 100
		for n > 0 {
			if n == 9 {
				s += "CM"
				break
			} else if n >= 5 {
				s += "D"
				n -= 5
			} else if n >= 4 {
				s += "CD"
				break
			} else {
				for j := 0; j < n; j++ {
					s += "C"
				}
				break
			}
		}
		num = num % 100
	}
	if num >= 10 {
		n := num / 10
		for n > 0 {
			if n == 9 {
				s += "XC"
				break
			} else if n >= 5 {
				s += "L"
				n -= 5
			} else if n >= 4 {
				s += "XL"
				break
			} else {
				for j := 0; j < n; j++ {
					s += "X"
				}
				break
			}
		}
		num = num % 10
	}
	for num > 0 {
		if num == 9 {
			s += "IX"
			break
		} else if num >= 5 {
			s += "V"
			num -= 5
		} else if num >= 4 {
			s += "IV"
			break
		} else {
			for j := 0; j < num; j++ {
				s += "I"
			}
			break
		}
	}

	return
}
