func longestBalanced(s string) int {
    m := 0
    for i := range len(s) {
        if len(s) - i < m { break }
        counter := [26]int{}
        maxCount := 0
        noChar := 0
        noMaxChar := 0
        for j:=i; j < len(s); j++ {
            idx := s[j] - 'a'
            counter[idx]++
            count := counter[idx]
            if count == 1 { noChar++ }
            if count == maxCount { noMaxChar++ }
            if count > maxCount {
                maxCount = count
                noMaxChar = 1
            }
            if noChar == noMaxChar && m < j-i+1 {
                m = j-i+1
            }
        }
    }
    return m    
}