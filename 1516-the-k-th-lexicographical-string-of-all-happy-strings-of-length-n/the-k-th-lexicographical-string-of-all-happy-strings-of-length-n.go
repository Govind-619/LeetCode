func generateString(n int, current string, count *int, k int, result *string) {
    if *count == k {
        return
    }

    if len(current) == n {
        *count++
        if *count == k {
            *result = current
        }
        return
    }

    for _, ch := range []rune{'a', 'b', 'c'} {
        if len(current) > 0 && current[len(current)-1] == byte(ch) {
            continue
        }
        generateString(n, current+string(ch), count, k, result)
    }
}

func getHappyString(n int, k int) string {
    count := 0
    result := ""
    generateString(n, "", &count, k, &result)
    
    if count < k {
        return ""
    }
    return result
}