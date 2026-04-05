func judgeCircle(moves string) bool {
    x,y:= 0,0
    for _,v:= range moves{
        switch v{
            case 'R':
            y++
            case 'L':
            y--
            case 'U':
            x++
            case 'D':
            x--
        }
    }
    return x==0 && y==0
}