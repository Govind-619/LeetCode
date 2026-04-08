func xorAfterQueries(nums []int, queries [][]int) int {
    const MOD int64 = 1000000007

    for _, q := range queries {
        l, r, k, v := q[0], q[1], q[2], q[3]

        for i := l; i <= r; i += k {
            nums[i] = int((int64(nums[i]) * int64(v)) % MOD)
        }
    }

    ans := 0
    for _, num := range nums {
        ans ^= num
    }

    return ans
}