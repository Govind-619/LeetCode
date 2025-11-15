func numberOfSubstrings(s string) (count int) {
	n := len(s)
	pre := make([]int, n+1)
	pre[0] = -1
	for i := 0; i < n; i++ {
		if i == 0 || i > 0 && s[i-1] == '0' {
			pre[i+1] = i
		} else {
			pre[i+1] = pre[i]
		}
	}

	for i := 1; i <= n; i++ {
		var cnt0 int
		if s[i-1] == '0' {
			cnt0 = 1
		}

		for j := i; j > 0 && cnt0*cnt0 <= n; j, cnt0 = pre[j], cnt0 + 1 {
			cnt1 := (i - pre[j]) - cnt0
			if cnt0 * cnt0 <= cnt1 {
				count += min(j - pre[j], cnt1 - cnt0 * cnt0 + 1)
			}
		}
	}
    
	return
}