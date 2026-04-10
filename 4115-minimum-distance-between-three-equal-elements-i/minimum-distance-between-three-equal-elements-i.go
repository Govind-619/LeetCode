func minimumDistance(nums []int) int {
    ans := -1
    for i := range nums {
        for j := range nums {
            for k := range nums {
                if i == j || j == k || i == k {
                    continue
                }

                if nums[i] != nums[j] || nums[j] != nums[k] || nums[i] != nums[k] {
                    continue
                }

                bruh := []int{i, j, k}
                sort.Ints(bruh)
                dist := (bruh[2] - bruh[0]) * 2 // the brain on this kid
                if ans == -1 || dist < ans {
                    ans = dist
                }
            }
        }
    }
    return ans
}