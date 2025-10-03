func passThePillow(n int, time int) int {
    cycle := 2*n - 2        // forward + backward cycle
    t := time % cycle        // effective seconds
    if t < n {
        return t + 1        // +1 because labels start from 1
    }
    return 2*n - 1 - t      // moving backward
}
 