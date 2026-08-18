func largestInteger(nums []int, k int) int {
	n := len(nums)

	freq := [51]int{}

	for _, x := range nums {
		freq[x]++
	}

	if k == 1 {
		for x := 50; x >= 0; x-- {
			if freq[x] == 1 {
				return x
			}
		}

		return -1
	}

	if k == n {
		answer := 0

		for _, x := range nums {
			if x > answer {
				answer = x
			}
		}

		return answer
	}

	answer := -1

	if freq[nums[0]] == 1 && nums[0] > answer {
		answer = nums[0]
	}

	if freq[nums[n-1]] == 1 && nums[n-1] > answer {
		answer = nums[n-1]
	}

	return answer
}