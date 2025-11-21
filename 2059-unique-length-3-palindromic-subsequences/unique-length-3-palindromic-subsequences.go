import(
"strings"
)

func countPalindromicSubsequence(s string) int {
    result := make(map[string]bool)
    letters := make(map[rune]bool)
    
    // Collect unique characters
    for _, c := range s {
        letters[c] = true
    }
    
    for char := range letters {
        first := strings.IndexRune(s, char)
        last := strings.LastIndexFunc(s, func(r rune) bool {
		return r == char
	    })
        
        if first < last {
            // Get all unique characters between first and last
            midChars := make(map[rune]bool)
            for i := first + 1; i < last; i++ {
                midChars[rune(s[i])] = true
            }
            
            for mid := range midChars {
                palindrome := string(char) + string(mid) + string(char)
                result[palindrome] = true
            }
        }
    }
    
    return len(result)
}