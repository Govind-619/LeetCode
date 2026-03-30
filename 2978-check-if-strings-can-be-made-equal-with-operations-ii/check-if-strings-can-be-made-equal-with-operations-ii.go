func checkStrings(s1 string, s2 string) bool {
    even := make([]int, 26)
    odd := make([]int, 26)

    for i := 0; i < len(s1); i++ {
        if i%2 == 0 {
            even[s1[i]-'a']++
            even[s2[i]-'a']--
        } else {
            odd[s1[i]-'a']++
            odd[s2[i]-'a']--
        }
    }

    for i := 0; i < 26; i++ {
        if even[i] != 0 || odd[i] != 0 {
            return false
        }
    }

    return true
}