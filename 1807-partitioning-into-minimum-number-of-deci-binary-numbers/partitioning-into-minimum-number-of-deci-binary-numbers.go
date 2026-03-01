func minPartitions(n string) int {
        max := 0
        for _, c := range n{
            if val := int(c - '0') 
            val == 9 {
                return 9
            } else if val > max {
                max = val
            }
        }
        return max
}