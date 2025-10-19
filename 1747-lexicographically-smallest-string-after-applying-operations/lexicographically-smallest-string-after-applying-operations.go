func findLexSmallestString(s string, a int, b int) string {
	rotate := func(s string) string {
		return s[b:] + s[:b]
	}
	add := func(s string) string {
		arr := []byte(s)
		for i := 1; i < len(arr); i += 2 {
			arr[i] = (((arr[i] - '0') + byte(a)) % 10) + '0'
		}
		return string(arr)
	}

	q := []string{s} // queue
	m := make(map[string]bool) // hash-table
	m[s] = true
	i := 0
	for i < len(q) {
        // q[i] is the next unprocessed element from queue
		s := rotate(q[i]) // rotate it
		if _, ok := m[s]; !ok { // place result to queue if we have not seen it
			q = append(q, s)
			m[s] = true
		}
		s = add(q[i]) // add to it
		if _, ok := m[s]; !ok { // place result to queue if we have not seen it
			q = append(q, s)
			m[s] = true
		}
		i++ // move to next unprocessed element
	}
	return slices.Min(q)
}