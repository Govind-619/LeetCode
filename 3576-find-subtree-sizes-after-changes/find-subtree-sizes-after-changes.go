func findSubtreeSizes(parent []int, s string) []int {
    n := len(parent)
    children := make([][]int, n)
    for x := 1; x < n; x++ {
        children[parent[x]] = append(children[parent[x]], x)
    }
    size, last := slices.Repeat([]int{1}, n), slices.Repeat([]int{-1}, 26)  // last[c]: the deepest node with letter c on the path from the root
    var dfs func(x int)
    dfs = func(x int) {
        c := s[x] - 'a'
        above := last[c]  // the closest ancestor with the same letter, or -1
        last[c] = x
        for _, child := range children[x] {
            dfs(child)
        }
        last[c] = above
        if x > 0 {  // the root has no parent
            to := parent[x]
            if above >= 0 {
                to = above
            }
            size[to] += size[x]
        }
    }
    dfs(0)
    return size
}