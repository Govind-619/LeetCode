func minJumps(arr []int) int {

	n := len(arr)

	if n == 1 {
		return 0
	}

	mp := make(map[int][]int)

	for i, val := range arr {
		mp[val] = append(mp[val], i)
	}

	queue := []int{0}

	visited := make([]bool, n)

	visited[0] = true

	steps := 0

	for len(queue) > 0 {

		size := len(queue)

		for i := 0; i < size; i++ {

			idx := queue[0]
			queue = queue[1:]

			if idx == n-1 {
				return steps
			}

			if idx-1 >= 0 && !visited[idx-1] {
				visited[idx-1] = true
				queue = append(queue, idx-1)
			}

			if idx+1 < n && !visited[idx+1] {
				visited[idx+1] = true
				queue = append(queue, idx+1)
			}

			for _, nextIdx := range mp[arr[idx]] {

				if !visited[nextIdx] {
					visited[nextIdx] = true
					queue = append(queue, nextIdx)
				}
			}

			mp[arr[idx]] = []int{}
		}

		steps++
	}

	return -1
}