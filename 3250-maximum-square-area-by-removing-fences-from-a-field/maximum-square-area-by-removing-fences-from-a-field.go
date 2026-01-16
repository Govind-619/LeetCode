func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	hFences = append([]int{1}, append(hFences, m)...)
	vFences = append([]int{1}, append(vFences, n)...)
	sort.Ints(hFences)
	sort.Ints(vFences)

	hDiffs := make(map[int]bool)
	for i := 0; i < len(hFences)-1; i++ {
		for j := i + 1; j < len(hFences); j++ {
			hDiffs[hFences[j]-hFences[i]] = true
		}
	}

	maxArea := int64(-1)
	for x := 0; x < len(vFences)-1; x++ {
		for y := x + 1; y < len(vFences); y++ {
			vDiff := vFences[y] - vFences[x]
			if hDiffs[vDiff] {
				maxArea = max(maxArea, int64(vDiff*vDiff))
			}
		}
	}

	return int(maxArea % 1_000_000_007)
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
