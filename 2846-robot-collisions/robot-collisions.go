
func survivedRobotsHealths(pos []int, hp []int, dir string) []int {
	n := len(pos)

	// Step 1: create indices [0..n-1]
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	// Step 2: sort indices based on position
	sort.Slice(indices, func(i, j int) bool {
		return pos[indices[i]] < pos[indices[j]]
	})

	// Stack pointer 
	nr := 0

	for i := 0; i < n; i++ {
		right := indices[i]
		isLeft := dir[right] == 'L'

		if isLeft && nr > 0 {
			// collision with last right-moving robot
			left := indices[nr-1]

			if hp[right] < hp[left] {
				// right robot dies
				hp[right] = 0
				hp[left]--
			} else if hp[right] > hp[left] {
				// left robot dies, right continues
				hp[left] = 0
				hp[right]--
				nr--
				i-- // reprocess current robot
			} else {
				// both die
				hp[right] = 0
				hp[left] = 0
				nr--
			}
		} else {
			// push into stack if moving right
			indices[nr] = indices[i]
			if !isLeft {
				nr++
			}
		}
	}

	// Collect surviving healths
	result := []int{}
	for i := 0; i < n; i++ {
		if hp[i] > 0 {
			result = append(result, hp[i])
		}
	}

	return result
}