func findDifferentBinaryString(nums []string) string {
    ans := ""
    for i, num := range nums{
        switch num[i]{
        case '0':
            ans = ans + "1"
        case '1':
            ans = ans + "0"
        default:
            panic("Invalid value should be 1 or 0 only")
        }
    }
    return ans
}