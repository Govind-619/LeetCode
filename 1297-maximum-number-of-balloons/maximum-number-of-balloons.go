func maxNumberOfBalloons(text string) int {
	r := make(map[rune]int)
	
	// Count frequencies of all characters in the text
	for _, v := range text {
		r[v]++
	}

	// Calculate how many times each character can fulfill its role in "balloon"
	b := r['b']
	a := r['a']
	l := r['l'] / 2
	o := r['o'] / 2
	n := r['n']

	// The result is the bottleneck (minimum) of all required characters
	return min(b, a, l, o, n)
}