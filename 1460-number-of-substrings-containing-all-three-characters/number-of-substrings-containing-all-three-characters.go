func numberOfSubstrings(s string) int {
    count := [3]int{}
    result, left := 0, 0

    for right, char := range s {
        count[char-'a']++

        for count[0] > 0 && count[1] > 0 && count[2] > 0 {
            result += len(s) - right
            count[s[left]-'a']--
            left++
        }
    }

    return result
}