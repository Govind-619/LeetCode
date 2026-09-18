import (
    "slices"
    "cmp"
)
func maxNumOfSubstrings(s string) []string {
    var first, last [26]int = [26]int{}, [26]int{};
    for i := 0; i < 26; i++ {
        first[i] = -1
        last[i] = -1
    }
    // Find first and last occurences of each character in a string
    // 26 possible substrings without validation
    for i, ch := range s {
        var ptr int = int(ch - 'a')
        if first[ptr] == -1 { first[ptr] = i }
        last[ptr] = i
    }
    // Find actual length of substring such start with i'th character
    var substrings [][2]int = [][2]int{}
    for i := 0; i < 26; i++ {
        if first[i] == -1 { continue }
        var start, end int = first[i], last[i]
        var needToStartAgain bool = true
        for needToStartAgain {
            needToStartAgain = false
            for j := start + 1; j < end; j++ {
                var curPtr int = int(s[j] - 'a')
                // If we meet character that firstly occur earlier that our start
                // Can we up to 26 rescanning at all
                if first[curPtr] < start {
                    start = first[curPtr]
                    needToStartAgain = true
                    break
                }
                end = max(end, last[curPtr])
            }
        }
        substrings = append(substrings, [2]int{start, end})
    }
    // Sort by earliest start and minimum length
    slices.SortFunc(substrings, func(x, y [2]int) int {
        if x[1] != y[1] { return cmp.Compare(x[1], y[1]) }
        return cmp.Compare(-x[0], -y[0])
    })
    // Use scheduler selecting events idea
    var ptr int = 0
    var output []string = []string{}
    for i := 0; i < len(substrings); i++ { // Up to 26 possible substrings
        if substrings[ptr][1] < substrings[i][0] {
            output = append(output, s[substrings[ptr][0]: substrings[ptr][1] + 1])
            ptr = i
        }
    }
    // string has always at least one substring(itself)
    output = append(output, s[substrings[ptr][0]: substrings[ptr][1] + 1])
    return output
}
