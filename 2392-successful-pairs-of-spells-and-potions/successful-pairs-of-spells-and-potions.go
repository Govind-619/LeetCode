import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)
    m := len(potions)
    res := make([]int, len(spells))

    for i, s := range spells {
        // Calculate the minimum potion strength required
        minPotion := float64(success) / float64(s)

        // Find the first potion >= minPotion
        idx := sort.Search(m, func(j int) bool {
            return float64(potions[j]) >= minPotion
        })

        // Count how many potions are successful
        res[i] = m - idx
    }
    return res
}
