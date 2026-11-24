func maxLength(nums []int) int {
    best, left, product := 2, 0, 1
    for right, x := range nums {
        for gcd(product, x) > 1 {
            product /= nums[left]
            left++
        }
        product *= x
        best = max(best, right-left+1)
    }
    return best
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}