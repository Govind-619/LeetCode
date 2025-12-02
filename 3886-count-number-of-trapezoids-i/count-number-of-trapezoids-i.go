func countTrapezoids(points [][]int) int {
    const MOD = 1000000007
    horLines := map[int]int{}
    for _, point := range points {
        horLines[point[1]]++
    }
    var total, pairs int
    for _, x := range horLines {
        total += pairs * (x*(x-1)/2)
        pairs += (x*(x-1)/2)
    }
    return total % MOD
}