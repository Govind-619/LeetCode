func maxDistance(side int, points [][]int, k int) int {
	n := len(points)
	pos := make([]int, n)

	// Map each point to a coordinate along the perimeter [0, 4*side)
	for i, p := range points {
		x, y := p[0], p[1]
		if y == 0 {
			pos[i] = x
		} else if x == side {
			pos[i] = side + y
		} else if y == side {
			pos[i] = 2*side + (side - x)
		} else { // x == 0
			pos[i] = 3*side + (side - y)
		}
	}

	sort.Ints(pos)

	L := 4 * side

	// Build extended array for circular wrap
	ext := make([]int, 2*n)
	copy(ext, pos)
	for i := 0; i < n; i++ {
		ext[n+i] = pos[i] + L
	}

	// Binary search helper (equivalent to bisect_left)
	bisectLeft := func(arr []int, target, lo, hi int) int {
		idx := sort.Search(hi-lo, func(i int) bool {
			return arr[lo+i] >= target
		})
		return lo + idx
	}

	var canPlace func(d int) bool
	canPlace = func(d int) bool {
		for start := 0; start < n; start++ {
			cur := start
			last := ext[start]
			valid := true
			limit := start + n

			for i := 1; i < k; i++ {
				target := last + d
				lo := bisectLeft(ext, target, cur+1, limit)

				if lo == limit {
					valid = false
					break
				}

				last = ext[lo]
				cur = lo
			}

			// Check wrap-around distance
			if valid && (ext[start]+L-last) >= d {
				return true
			}
		}
		return false
	}

	low, high := 0, 2*side
	for low < high {
		mid := (low + high + 1) / 2
		if canPlace(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return low
}