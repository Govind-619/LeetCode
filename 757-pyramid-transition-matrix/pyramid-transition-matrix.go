func pyramidTransition(b string, a []string) bool {
    mem := map[[2]string]bool{} // memory
    var dfs func(cr, nr string) bool // current row, next row (based on match from allowed)
    dfs = func(cr, nr string) bool {
        if _, ok := mem[[2]string{nr, cr}]; ok {
            return mem[[2]string{nr, cr}]
        } // check mem
        if len(cr) == 1 { return true } // if reached top row, return true
        crid := len(nr) // get current index of current row
        if crid >= len(cr)-1 { // if we are finished with current row, switch to next row
            return dfs(nr, "")
        }
        res := false
        for k := range a { // iterate allowed array
            if a[k][:2] == cr[crid:crid+2] {
                res = res || dfs(cr, nr+string(a[k][2])) // if found match, do next 'dfs'
            }
        }
        mem[[2]string{nr, cr}] = res // update mem
        return res
    }
    return dfs(b, "") // start dfs
}