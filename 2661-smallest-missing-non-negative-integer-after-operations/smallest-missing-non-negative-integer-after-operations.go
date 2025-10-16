func findSmallestInteger(nums []int, value int) int {
    // loop through nums, and put each number in its group.
    n := len(nums)
    m := make([]int, n)
    mod := func(x,y int) int { return ((x%y) + y) % y} // handles negative numbers as well
    for _,v := range nums {
        if curr := mod(v,value); curr < n {
            m[curr]++
        }
    }
    // see which numbers in the range 0..n are achieavble, we stop at the first one that we cannot achieve
    for nb := 0; nb < len(nums); nb++ {
        if m[nb%value] > 0 {
            m[nb%value]--
        } else {
            return nb
        }
    }
    return len(nums)
}