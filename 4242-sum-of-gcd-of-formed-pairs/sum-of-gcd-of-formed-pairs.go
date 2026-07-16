import "sort"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdSum(nums []int) int64 {

	n := len(nums)

	prefixGcd := make([]int, n)

	prefixMax := 0

	for i, x := range nums {
		if x > prefixMax {
			prefixMax = x
		}
		prefixGcd[i] = gcd(x, prefixMax)
	}

	sort.Ints(prefixGcd)

	var ans int64 = 0

	left := 0
	right := n - 1

	for left < right {
		ans += int64(gcd(prefixGcd[left], prefixGcd[right]))
		left++
		right--
	}

	return ans
}