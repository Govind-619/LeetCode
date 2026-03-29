func canBeEqual(s1 string, s2 string) bool {
    rn := []rune(s1)
    for i := 0; i <= 1; i++ {
        if s1[i] != s2[i] && s1[i+2] != s2[i+2] {
            rn[i], rn[i+2] = rn[i+2], rn[i]
        }
    }
    return string(rn) == s2
}