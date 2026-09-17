func minSumOfLengths(arr []int, target int) int {
	n := len(arr)

	const INF = int(^uint(0) >> 1)
	best := make([]int, n)

	for i := 0; i < n; i++ {
		best[i] = INF
	}

	left := 0
	sum := 0
	answer := INF

	for right := 0; right < n; right++ {
		sum += arr[right]

		for sum > target {
			sum -= arr[left]
			left++
		}

		if sum == target {
			currentLength := right - left + 1

			if left > 0 && best[left-1] != INF {
				candidate := currentLength + best[left-1]

				if candidate < answer {
					answer = candidate
				}
			}

			best[right] = currentLength
		}

		if right > 0 && best[right-1] < best[right] {
			best[right] = best[right-1]
		}
	}

	if answer == INF {
		return -1
	}

	return answer
}
