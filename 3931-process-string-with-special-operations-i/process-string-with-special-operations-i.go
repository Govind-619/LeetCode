import "unicode"

func processStr(s string) string {
    result:=[]rune{}
    for _,v:= range s{
        if unicode.IsLetter(v){
            result= append(result,v)
        }else if v=='#'{
            result= append(result,result...)
        }else if v=='%'{
            reverse(result)
        }else{
            l:= len(result)
            if l<2{
                result =[]rune{}
            }else{
                result= result[:l-1]
            }
        }
    }
    return string(result)
}

func reverse(s []rune){
    n:= len(s)
    for i,j:=0,n-1; i<j; i,j = i+1, j-1{
        s[i], s[j]= s[j], s[i]
    }
}