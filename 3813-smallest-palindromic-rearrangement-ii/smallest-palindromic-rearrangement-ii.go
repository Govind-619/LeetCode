func smallestPalindrome(s string, k int) string {
    freq := make([]int, 26)
    for _, c := range s {
        freq[c-'a']++
    }
    
    half := make([]int, 26)
    mid := ""
    m := 0
    for i := 0; i < 26; i++ {
        if freq[i] % 2 != 0 {
            mid += string(rune(i + 'a'))
        }
        half[i] = freq[i] / 2
        m += half[i]
    }
    
    getWays := func(f []int, targetK int) int {
        ways := 1
        currLen := 0
        for _, count := range f {
            if count > 0 {
                currLen += count
                n := currLen
                r := count
                
                if r > n - r {
                    r = n - r
                }
                curNCr := 1
                
                for i := 1; i <= r; i++ {
                    curNCr = curNCr * (n - i + 1) / i
                    if curNCr > targetK {
                        curNCr = targetK + 1
                        break
                    }
                }
                ways *= curNCr
                if ways > targetK {
                    return targetK + 1
                }
            }
        }
        return ways
    }
    
    if getWays(half, k) < k {
        return ""
    }
    
    firstHalf := []byte{}
    for i := 0; i < m; i++ {
        for c := 0; c < 26; c++ {
            if half[c] > 0 {
                half[c]--
                ways := getWays(half, k)
                
                if ways >= k {
                    firstHalf = append(firstHalf, byte(c+'a'))
                    break
                } else {
                    k -= ways
                    half[c]++
                }
            }
        }
    }
    
    res := string(firstHalf) + mid
    for i := len(firstHalf) - 1; i >= 0; i-- {
        res += string(firstHalf[i])
    }
    return res
}