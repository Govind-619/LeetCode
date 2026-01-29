func queryResults(limit int, queries [][]int) []int {
    colourPresent := []int{}

    f := make(map[int]int)   // ball -> color
    f1 := make(map[int]int)  // color -> count

    for _, v := range queries {
        ball := v[0]
        newColor := v[1]

        // If ball already has a color, decrement old color count
        if oldColor, exists := f[ball]; exists {
            f1[oldColor]--
            if f1[oldColor] == 0 {
                delete(f1, oldColor)
            }
        }

        // Assign new color
        f[ball] = newColor
        f1[newColor]++

        // Record number of distinct colors
        colourPresent = append(colourPresent, len(f1))
    }

    return colourPresent
}
