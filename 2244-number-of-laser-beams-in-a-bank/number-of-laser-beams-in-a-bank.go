func numberOfBeams(bank []string) int {
    res:= []int{}
    for _, v:= range bank{
        count:=0
        for _, ray:= range v{
            if ray=='1'{
                count++
            }
        }
        if count==0{
            continue
        }else{
            res= append(res, count)
        } 
    }
    output:=0
    for i:= 0; i<len(res)-1; i++{
        output += res[i]*res[i+1]
    }
    return output
}