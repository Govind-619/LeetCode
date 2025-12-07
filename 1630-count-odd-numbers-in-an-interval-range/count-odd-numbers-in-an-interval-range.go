func countOdds(low int, high int) int {
    length := high - low +1

    if length % 2 != 0 && low%2 != 0{
        return length/2+1
    }
    return length/2
}