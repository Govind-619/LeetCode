func maxMatrixSum(matrix [][]int) int64 {
	var total int64 = 0
	negative := 0
	minAbs := int64(math.MaxInt64)

	for _, row := range matrix {
		for _, cell := range row {
			absVal := int64(abs(cell))
			total += absVal

			if cell < 0 {
				negative++
			}

			if absVal < minAbs {
				minAbs = absVal
			}
		}
	}

	if negative%2 != 0 {
		total -= 2 * minAbs
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
