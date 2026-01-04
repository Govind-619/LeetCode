func sumFourDivisors(nums []int) int {
    totalSum := 0

    for _, n := range nums {
        count := 0
        sum := 0

        for i := 1; i*i <= n; i++ {
            if n%i == 0 {
                d1 := i
                d2 := n / i

                if d1 == d2 {
                    count++
                    sum += d1
                } else {
                    count += 2
                    sum += d1 + d2
                }

                if count > 4 {
                    break
                }
            }
        }

        if count == 4 {
            totalSum += sum
        }
    }

    return totalSum
}
