func minimumAbsDifference(arr []int) [][]int {
    if len(arr)<2{
        return [][]int{arr}
    }
    sort.Ints(arr)
    min:= abs(arr[1]-arr[0])
    for i:=2;i<len(arr); i++{
        if abs(arr[i]-arr[i-1])<min{
            min=abs(arr[i-1]-arr[i])
        }
    }
    ans:= make([][]int,0)
    for i:=1;i<len(arr); i++{
        if arr[i]-min==arr[i-1]{
            ans= append(ans,[]int{arr[i-1],arr[i]})
        }
    }
    return ans
}

func abs(x int)int{
    if x<0{
        return -x
    }
    return x
}