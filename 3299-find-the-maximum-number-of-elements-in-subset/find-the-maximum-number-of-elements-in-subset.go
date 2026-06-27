func maximumLength(nums []int) int {
	freq := make(map[int64]int)

	for _, x := range nums {
		freq[int64(x)]++
	}

	ans := 1

	if cnt, ok := freq[1]; ok {
		if cnt%2 == 1 {
			if cnt > ans {
				ans = cnt
			}
		} else {
			if cnt-1 > ans {
				ans = cnt - 1
			}
		}
	}

	for start := range freq {
		if start == 1 {
			continue
		}

		cur := start
		length := 0

		for {
			cnt, ok := freq[cur]
			if !ok {
				break
			}

			if cnt >= 2 {
				length += 2

				cur = cur * cur
			} else {
				length++
				break
			}
		}

		if length%2 == 0 {
			length--
		}

		if length > ans {
			ans = length
		}
	}

	return ans
}