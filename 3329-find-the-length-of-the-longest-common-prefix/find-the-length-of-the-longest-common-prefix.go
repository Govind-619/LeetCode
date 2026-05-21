func longestCommonPrefix(arr1 []int, arr2 []int) int {
    
    prefixes := make(map[int]bool)

    for _, num := range arr1 {

        x := num

        for x > 0 {

            prefixes[x] = true

            x /= 10
        }
    }

    ans := 0

    for _, num := range arr2 {

        x := num

        for x > 0 {

            if prefixes[x] {

                length := 0
                temp := x

                for temp > 0 {
                    length++
                    temp /= 10
                }

                if length > ans {
                    ans = length
                }

                break
            }

            x /= 10
        }
    }

    return ans
}