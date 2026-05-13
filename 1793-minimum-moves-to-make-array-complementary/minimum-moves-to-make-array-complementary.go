func minMoves(nums []int, limit int) int {
    
    n := len(nums)

    diff := make([]int, 2*limit+2)

    for i := 0; i < n/2; i++ {

        a := nums[i]
        b := nums[n-1-i]

        if a > b {
            a, b = b, a
        }

        diff[a+1] -= 1
        diff[b+limit+1] += 1

        diff[a+b] -= 1
        diff[a+b+1] += 1
    }

    pairs := n / 2

    current := pairs * 2

    answer := current

    for sum := 2; sum <= 2*limit; sum++ {

        current += diff[sum]

        if current < answer {
            answer = current
        }
    }

    return answer
}