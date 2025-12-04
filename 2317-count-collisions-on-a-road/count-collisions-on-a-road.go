func countCollisions(directions string) int {
    r := []rune(directions)
    n := len(r)

    left := 0
    for left < n && r[left] == 'L' {
        left++
    }

    right := n - 1
    for right >= 0 && r[right] == 'R' {
        right--
    }

    collisions := 0
    for i := left; i <= right; i++ {
        if r[i] != 'S' {
            collisions++
        }
    }

    return collisions
}
