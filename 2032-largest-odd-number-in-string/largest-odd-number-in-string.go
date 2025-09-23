func largestOddNumber(num string) string {
    for i := len(num) - 1; i >= 0; i-- {
        d := num[i]
        if (d-'0')%2 == 1 {           // check if digit is odd
            return num[:i+1]
        }
    }
    return ""
}
