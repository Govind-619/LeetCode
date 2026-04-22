func twoEditWords(queries []string, dictionary []string) []string {
    result:= []string{}
    for _,query:= range queries{
        for _,word:= range dictionary{
            if sameWord(query,word){
                result= append(result,query)
                break
            }
        }
    }
    return result
}

func sameWord(s1, s2 string)bool{
    if len(s1)!=len(s2){
        return false
    }
    diff:=0
    for i:=0; i<len(s1); i++{
        if s1[i]!=s2[i]{
            diff++
        }
    }
    if diff<=2{
        return true
    }
    return false
}