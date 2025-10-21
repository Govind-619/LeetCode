func maxFrequency(nums []int, k int, numOperations int) int {
	MAX := 100001
	freq := make([]int, MAX+1)
	inc := make([]int, MAX+1)
	l := MAX
	r := 0
	for _, num := range nums {
		tl := max(0, num-k)
		tr := min(MAX, num+k+1)
		inc[tl]++
		inc[tr]--
		freq[num]++

		l = min(l, tl)
		r = max(r, tr)
	}

	ans := 0
	curr := 0
	for i := l; i <= r; i++ {
		curr += inc[i]
		x := min(numOperations, curr-freq[i]) + freq[i]
		ans = max(ans, x)
	}

	return ans
}