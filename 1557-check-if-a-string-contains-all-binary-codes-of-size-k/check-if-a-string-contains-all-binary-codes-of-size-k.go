func hasAllCodes(s string, k int) bool {
    codeMap := make(map[string]bool)
    sLen := len(s)
    maxSize := 1 << k

    for i := 0; i <= sLen - k && len(codeMap) < maxSize; i++ {
        codeMap[s[i:i + k]] = true
    }

    return len(codeMap) == maxSize
}