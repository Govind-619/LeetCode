import "sort"

func minimumCost(cst []int) int {
    sort.Slice(cst, func(i, j int) bool {
        return cst[i] > cst[j]
    })

    ans := 0

    for i, val := range cst {
        if i%3 != 2 {
            ans += val
        }
    }

    return ans
}