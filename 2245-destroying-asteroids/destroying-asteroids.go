import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids)

    cur := int64(mass)

    for _, a := range asteroids {
        if cur < int64(a) {
            return false
        }
        cur += int64(a)
    }

    return true
}