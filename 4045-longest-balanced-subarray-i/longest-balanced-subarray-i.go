func longestBalanced(nums []int) int {
    n := len(nums)
    ans := 0

    for i := 0; i < n; i++ {
        evenFreq := make(map[int]int)
        oddFreq := make(map[int]int)

        for j := i; j < n; j++ {
            if nums[j]%2 == 0 {
                evenFreq[nums[j]]++
            } else {
                oddFreq[nums[j]]++
            }

            if len(evenFreq) == len(oddFreq) {
                if j-i+1 > ans {
                    ans = j - i + 1
                }
            }
        }
    }
    return ans
}
