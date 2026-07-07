func sumAndMultiply(n int) int64 {
    var x int64 = 0

    var sum int64 = 0

    divisor := 1
    for n/divisor >= 10 {
        divisor *= 10
    }

    for divisor > 0 {
        digit := n / divisor

        n %= divisor

        if digit != 0 {
            x = x*10 + int64(digit)

            sum += int64(digit)
        }

        divisor /= 10
    }

    return x * sum
}