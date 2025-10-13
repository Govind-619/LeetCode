func removeAnagrams(words []string) []string {
    res := []string{words[0]}
    for i := 1; i < len(words); i++ {
        prev := []rune(words[i-1])
        curr := []rune(words[i])
        sort.Slice(prev, func(i, j int) bool { return prev[i] < prev[j] })
        sort.Slice(curr, func(i, j int) bool { return curr[i] < curr[j] })
        if string(prev) != string(curr) {
            res = append(res, words[i])
        }
    }
    return res
}
