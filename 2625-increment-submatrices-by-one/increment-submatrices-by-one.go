func rangeAddQueries(n int, queries [][]int) [][]int {
    arr:= make([][]int,n)
    for i:=0; i<n; i++{
        arr[i]= make([]int,n)
    }
    for _,v:= range queries{
        for i:=v[0]; i<=v[2]; i++{
            for j:=v[1]; j<=v[3]; j++{
                arr[i][j]++
            }
        }
    }
    return arr
}