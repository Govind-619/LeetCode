func minOperations(nums []int) int {
    ans := 0
    // stack will hold a non-decreasing sequence (with a leading 0)
    stack := make([]int, 0, len(nums)+1)
    stack = append(stack, 0)

    for _, num := range nums {
        // pop while stack top > current num
        for len(stack) > 0 && stack[len(stack)-1] > num {
            stack = stack[:len(stack)-1]
        }
        // if stack top < current num, we have a new segment -> increment operations
        if stack[len(stack)-1] < num {
            ans++
            stack = append(stack, num)
        }
        // if stack top == num, do nothing (we are continuing an existing segment)
    }

    return ans
}
