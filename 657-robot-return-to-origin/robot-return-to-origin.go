func judgeCircle(moves string) bool {
    Ox,Oy:=0,0
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
    return Ox==x && Oy==y
}