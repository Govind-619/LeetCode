func minMirrorPairDistance(nums []int) int {
    // prev stores the most recent index i where reverse(nums[i]) = key
    prev := make(map[int]int)
    n := len(nums)
    ans := n + 1

    for i, x := range nums {
        // 1. Check if the current value x has appeared as a reverse of a previous element
        if j, ok := prev[x]; ok {
            if i - j < ans {
                ans = i - j
            }
        }

        // 2. Store the reverse of the current number mapping to the current index.
        // This allows future elements to find this index if they match the reverse.
        prev[reverse(x)] = i
    }

    if ans > n {
        return -1
    }
    return ans
}

func reverse(x int) int {
    // Handle 0 or negative if necessary, though typical for this problem x > 0
    res := 0
    for x > 0 {
        res = res*10 + x%10
        x /= 10
    }
    return res
}