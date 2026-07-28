func smallestPalindrome(s string) string {
    // frequency count
    freq := [26]int{}
    for _, ch := range s {
        freq[ch-'a']++
    }

    // build left half
    left := []rune{}
    var middle rune
    for i := 0; i < 26; i++ {
        if freq[i]%2 == 1 {
            middle = rune('a' + i) // odd count char for the middle (at most one)
        }
        for j := 0; j < freq[i]/2; j++ {
            left = append(left, rune('a'+i))
        }
    }

    // build result: left + middle (if any) + reversed(left)
    res := []rune{}
    res = append(res, left...)
    if middle != 0 {
        res = append(res, middle)
    }
    for i := len(left) - 1; i >= 0; i-- {
        res = append(res, left[i])
    }
    return string(res)
}