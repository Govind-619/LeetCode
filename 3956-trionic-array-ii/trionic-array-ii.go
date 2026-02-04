func maxSumTrionic(nums []int) int64 {
	const mark = math.MinInt / 2

	pBeg, pInc1, pDec, pInc2 := nums[0], mark, mark, mark
	result := math.MinInt
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if pBeg < n {
			pBeg, pInc1, pDec, pInc2 = n, max(pInc1+n, pBeg+n), mark, max(pInc2+n, pDec+n)
			result = max(result, pInc2)
		} else if n < pBeg {
			pBeg, pInc1, pDec, pInc2 = n, mark, max(pDec+n, pInc1+n), mark
		} else {
			pBeg, pInc1, pDec, pInc2 = n, mark, mark, mark
		}
	}

	return int64(result)
}