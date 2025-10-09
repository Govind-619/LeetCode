func minTime(skill []int, mana []int) int64 {
    fin := make([]int64, len(skill))
    for _,m := range mana {
        fin[0] = fin[0] + int64(skill[0] * m)
        dif := int64(0)
        if len(skill) > 1 { dif = max(dif, fin[1]-fin[0]) }
        for i:=1; i<len(skill); i++ {
            fin[i] = fin[i-1] + int64(skill[i] * m)
            if i < len(skill)-1 {
                if i < len(skill)-1 && fin[i+1]-fin[i] > dif { dif = fin[i+1]-fin[i] }
            }
        }
        if dif > 0 {
            for i,f := range fin {
                fin[i] = f + dif
            }
        }
    }
    return fin[len(skill)-1]
}