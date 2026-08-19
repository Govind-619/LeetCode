func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	rows := map[int]int{}

	m1 :=     0b111100
	m2 :=   0b11110000
	m3 := 0b1111000000

	for _, seat := range reservedSeats {
		i := seat[0]
		j := seat[1]

		rows[i] |= 1 << j
	}

	count := 0

	for _, row := range rows {
		if row&m1 == 0 {
			count++
			if row&m3 == 0 {
				count++
			}
		} else if row&m2 == 0 {
			count++
		} else if row&m3 == 0 {
			count++
		}
	}

	count += (n - len(rows)) * 2
	return count
}