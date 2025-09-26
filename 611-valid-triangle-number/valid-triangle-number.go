import "sort"

func triangleNumber(nums []int) int {
    sort.Ints(nums)
    count := 0

    for k := len(nums) - 1; k >= 2; k-- {
        i, j := 0, k-1
        for i < j {
            if nums[i]+nums[j] > nums[k] {
                count += j - i   // all pairs (i...j-1, j) are valid
                j--
            } else {
                i++
            }
        }
    }

    return count
}
