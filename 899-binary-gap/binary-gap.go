func binaryGap(n int) int {
    s:= fmt.Sprintf("%b", n)
    arr:=[]int{}
    for i,v:= range s{
        if v=='1'{
            arr= append(arr,i)
        }
    }
    value:=0

    for i:=1; i<len(arr); i++ {
        if arr[i]-arr[i-1]>value{
            value= arr[i]-arr[i-1]
        }
    }
    return value
}