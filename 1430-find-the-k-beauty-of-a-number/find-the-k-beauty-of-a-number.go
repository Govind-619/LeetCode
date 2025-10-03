func divisorSubstrings(num int, k int) int {
	count := 0
	s := strconv.Itoa(num)

	for i := 0; i <= len(s)-k; i++ {
		sub := s[i : i+k] // take k-length substring
		n, _ := strconv.Atoi(sub)
		if n != 0 && num%n == 0 { // avoid divide by zero
			count++
		}
	}
	return count
}