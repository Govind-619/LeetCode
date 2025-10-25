func totalMoney(n int) int {
    weeks := n / 7
    days := n % 7

    sum := 0

    // Sum for complete weeks
    // Each week starts higher than the previous one
    for w := 0; w < weeks; w++ {
        // Week w starts at (1+w), goes up to (7+w)
        sum += (7 * (2*(1+w) + 6)) / 2  // arithmetic progression sum
    }

    // Sum for remaining days in the last (incomplete) week
    for d := 0; d < days; d++ {
        sum += weeks + 1 + d
    }

    return sum
}
