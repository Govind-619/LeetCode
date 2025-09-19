func maxProduct(n int) int {
    max1, max2 := 0, 0
	for n > 0 {
		d := n % 10
		if d >= max1 {
			max2 = max1
			max1 = d
		} else if d > max2 {
			max2 = d
		}
		n /= 10
	}
	return max1 * max2
}