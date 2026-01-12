func minTimeToVisitAllPoints(points [][]int) int {
    count:=0
    for i:=0; i<len(points)-1; i++{
       x1:= points[i][0]
       y1:= points[i][1]
       x2:= points[i+1][0]
       y2:= points[i+1][1]
       count += max(abs(x1-x2), abs(y1-y2))
    }
    return count
}

func max(x,y int)int{
    if x>y{
        return x
    }
    return y
}

func abs(x int)int{
 if x<0{
    return -x
 }
 return x
}