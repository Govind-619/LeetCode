func maxArea(height []int) int {
	res := -1
	l, r := 0, len(height)-1

	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > res {
			res = area
		}

		if height[l] < height[r] {
			l++
		} else if height[l] > height[r] {
			r--
		} else {
			r--
			l++
		}
	}

	return res
}