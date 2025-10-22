func areSimilar(mat [][]int, k int) bool {
    k %= len(mat[0])
    for _,v := range mat {
        k = -k
        for j,w := range v {
            a := j + k
            if a < 0 {
                a += len(v)
            } else if a >= len(v) {
                a -= len(v)
            }
            if w != v[a] {return false}
        }
    }
    return true
}