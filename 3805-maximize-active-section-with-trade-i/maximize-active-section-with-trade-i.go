func maxActiveSectionsAfterTrade(s string) int {
	ones := 0
	for _, ch := range s {
		if ch == '1' {
			ones++
		}
	}

	t := "1" + s + "1"

	type Run struct {
		ch  byte
		len int
	}

	runs := []Run{}

	for i := 0; i < len(t); i++ {
		if len(runs) == 0 || runs[len(runs)-1].ch != t[i] {
			runs = append(runs, Run{t[i], 1})
		} else {
			runs[len(runs)-1].len++
		}
	}

	best := 0

	for i := 1; i+1 < len(runs); i++ {
		if runs[i].ch == '1' &&
			runs[i-1].ch == '0' &&
			runs[i+1].ch == '0' {

			gain := runs[i-1].len + runs[i+1].len
			if gain > best {
				best = gain
			}
		}
	}

	return ones + best
}