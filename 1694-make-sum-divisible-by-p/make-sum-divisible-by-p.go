func minSubarray(nums []int, p int) int {
    total := 0
    for _,v := range nums {
        total+=v
    }
    targetRemainder := total % p
    if targetRemainder == 0 {
        return 0
    }

    // run the algorithm
    remainderToIdx := map[int]int{0:-1}
    runningSum := 0
    length := len(nums)+1
    for i := range nums {
        runningSum+=nums[i]
        runningSum%= p
        need := (runningSum - targetRemainder + p) % p
        if idx,ok := remainderToIdx[need]; ok {
            length = min(length, i-idx)
        }
        remainderToIdx[runningSum] = i
    }

    if length >= len(nums) {
        return -1
    }
    return length
}