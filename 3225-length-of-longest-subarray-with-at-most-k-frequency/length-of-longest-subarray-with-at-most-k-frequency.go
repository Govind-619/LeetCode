func maxSubarrayLength(nums []int, k int) int {
    freq := make(map[int]int)
    maxLen := 0
    i := 0

    for j := 0; j < len(nums); j++ {
        freq[nums[j]]++

        for freq[nums[j]] > k {
            freq[nums[i]]--
            i++
        }

        maxLen = max(maxLen, j-i+1)
    }

    return maxLen
}