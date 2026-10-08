func removeOuterParentheses(s string) string {
    l,r:=0,0
    flag:=0
    a:= []rune{}
    if s == ""{
        return s
    }else{
    for i:=0; i<len(s); i++{
        if s[i]=='('{
            flag++   
        }else{
            flag--
        }
        if flag == 0{
            r=i
            for k:=l+1; k<r; k++{
                a= append(a,rune(s[k]))
            }
            l=i+1
        }
    }  
    return string(a)
    }
}