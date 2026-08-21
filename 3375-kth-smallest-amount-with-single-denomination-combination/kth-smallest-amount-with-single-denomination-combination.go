func findKthSmallest(coins []int, k int) int64 {
	sort.Ints(coins)

	useful := make([]int64, 0)

	for _, coin := range coins {
		redundant := false

		for _, prev := range useful {
			if int64(coin)%prev == 0 {
				redundant = true
				break
			}
		}

		if !redundant {
			useful = append(useful, int64(coin))
		}
	}

	m := len(useful)

	low := int64(1)
	high := useful[0] * int64(k)

	totalMasks := 1 << m

	lcms := make([]int64, totalMasks)

	signs := make([]int64, totalMasks)

	for mask := 1; mask < totalMasks; mask++ {
		currentLCM := int64(1)
		bits := 0

		for i := 0; i < m; i++ {
			if mask&(1<<i) != 0 {
				g := gcd(currentLCM, useful[i])

				currentLCM /= g

				if currentLCM > high/useful[i] {
					currentLCM = high + 1
					break
				}

				currentLCM *= useful[i]
				bits++
			}
		}

		lcms[mask] = currentLCM

		if bits%2 == 1 {
			signs[mask] = 1
		} else {
			signs[mask] = -1
		}
	}

	count := func(x int64) int64 {
		result := int64(0)

		for mask := 1; mask < totalMasks; mask++ {
			if lcms[mask] <= x {
				result += signs[mask] * (x / lcms[mask])
			}
		}

		return result
	}

	for low < high {
		mid := low + (high-low)/2

		if count(mid) >= int64(k) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}