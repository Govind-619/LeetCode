func minimumPushes(word string) int {
    // count times
    m := make([]int, 26, 26) // Approach 3. letter->times
    for _, v := range word {
        m[v-97]++
    }
    
    // handle most used button of times
    sort.Slice(m, func(i, j int)bool { // Approach 2.
        return m[i] > m[j]
    })
    
    // cal
    ans := 0
    for i:=0; i<len(m); i++ {
        addCount := i/8 +1 // Intuition 1.
        if m[i] > 0 {
            ans+= addCount 
        }
    }
    return ans
}