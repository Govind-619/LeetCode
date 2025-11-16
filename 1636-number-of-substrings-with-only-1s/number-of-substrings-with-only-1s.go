func numSub(s string) int {
    res,currOnes,modulo := 0,0,int(1e9 +7)
    for i := range s {
        if s[i] == '1' {
            currOnes++
            res= (res+currOnes)%modulo
        } else {
            currOnes = 0
        }
    }
    return res
}