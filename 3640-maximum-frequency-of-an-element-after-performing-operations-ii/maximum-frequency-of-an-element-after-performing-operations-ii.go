type Pair struct {
	first  int64
	second int64
}

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	parr := make([]Pair, 0)
	for i := 0; i < len(nums); i++ {
		j := i
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		cnt := int64(j - i)
		parr = append(parr, Pair{int64(nums[i]), cnt})
		i = j - 1
	}

	prevCand := make([]Pair, 0)
	for i := 0; i < len(parr); i++ {
		prevCand = append(prevCand, parr[i])
		prevCand = append(prevCand, Pair{parr[i].first + int64(k), 0})
	}

	sort.Slice(prevCand, func(i, j int) bool {
		if prevCand[i].first == prevCand[j].first {
			return prevCand[i].second < prevCand[j].second
		}
		return prevCand[i].first < prevCand[j].first
	})

	candidates := make([]Pair, 0)
	for i := 0; i < len(prevCand); i++ {
		if i+1 < len(prevCand) && prevCand[i].first == prevCand[i+1].first {
			candidates = append(candidates, Pair{prevCand[i].first, prevCand[i].second + prevCand[i+1].second})
			i++
		} else {
			candidates = append(candidates, prevCand[i])
		}
	}

	ans := 0
	for i := 0; i < len(candidates); i++ {
		curr := candidates[i].first
		cnt := candidates[i].second

		l := sort.Search(len(nums), func(i int) bool {
			return int64(nums[i]) >= curr-int64(k)
		})
		r := sort.Search(len(nums), func(i int) bool {
			return int64(nums[i]) > curr+int64(k)
		})

		x := min(r-l-int(cnt), numOperations) + int(cnt)
		ans = max(ans, x)
	}

	return ans
}