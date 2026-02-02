func minimumCost(nums []int, k int, dist int) int64 {
	n := len(nums)
	targetK := k - 1
	temp := make([]int, n)
	copy(temp, nums)
	sort.Ints(temp)
	m := 0
	for i := 0; i < n; i++ {
		if i == 0 || temp[i] != temp[i-1] {
			temp[m] = temp[i]
			m++
		}
	}
	sorted := temp[:m]
	bitSum := make([]int64, m+1)
	bitCount := make([]int, m+1)

	update := func(r, v, c int) {
		for ; r <= m; r += r & -r {
			bitSum[r] += int64(v)
			bitCount[r] += c
		}
	}

	maxP := 1
	for maxP<<1 <= m {
		maxP <<= 1
	}
	var minExtra int64 = math.MaxInt64

	for i := 1; i < n; i++ {
		r := sort.SearchInts(sorted, nums[i]) + 1
		update(r, nums[i], 1)
		if i > dist+1 {
			oldV := nums[i-dist-1]
			oldR := sort.SearchInts(sorted, oldV) + 1
			update(oldR, -oldV, -1)
		}
		if i >= targetK {
			idx, cc, cs := 0, 0, int64(0)
			for p := maxP; p > 0; p >>= 1 {
				if idx+p <= m && cc+bitCount[idx+p] < targetK {
					idx += p
					cc += bitCount[idx]
					cs += bitSum[idx]
				}
			}
			if cc < targetK {
				cs += int64(targetK-cc) * int64(sorted[idx])
			}
			if cs < minExtra {
				minExtra = cs
			}
		}
	}
	return int64(nums[0]) + minExtra
}