func readBinaryWatch(turnedOn int) []string {
    var result []string

    for h := 0; h < 12; h++ {
        for m := 0; m < 60; m++ {
            if countBits(h)+countBits(m) == turnedOn {
                result = append(result, formatTime(h, m))
            }
        }
    }

    return result
}

func countBits(n int) int {
    count := 0
    for n > 0 {
        count += n & 1
        n >>= 1
    }
    return count
}

func formatTime(h, m int) string {
    if m < 10 {
        return fmt.Sprintf("%d:0%d", h, m)
    }
    return fmt.Sprintf("%d:%d", h, m)
}
