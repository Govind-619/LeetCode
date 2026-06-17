func processStr(s string, k int64) byte {
	n := len(s)

	lengths := make([]int64, n)
	var curLen int64 = 0

	for i := 0; i < n; i++ {
		c := s[i]

		if c >= 'a' && c <= 'z' {
			curLen++
		} else if c == '*' {
			if curLen > 0 {
				curLen--
			}
		} else if c == '#' {
			curLen *= 2
		} else {
		}

		lengths[i] = curLen
	}

	if k >= curLen {
		return '.'
	}

	for i := n - 1; i >= 0; i-- {
		c := s[i]

		var before int64
		if i > 0 {
			before = lengths[i-1]
		}

		if c >= 'a' && c <= 'z' {
			if k == before {
				return c
			}
		} else if c == '#' {
			if before > 0 {
				k %= before
			}
		} else if c == '%' {
			k = before - 1 - k
		} else {
		}
	}

	return '.'
}