func maxDistance(colors []int) int {
    n:= len(colors)
    max:=0
    for i:=0; i<n;i++{
        for j:=n-1; j>i; j--{
            if colors[i]!=colors[j]{
                max= j-i
            }
        }
    }
    for j:=n-1; j>=0; j--{
        for i:=0; i<j;i++{
            if colors[i]!=colors[j]{
                if j-i>max{
                     max= j-i
                }
            }
        }
    }
    return max
}