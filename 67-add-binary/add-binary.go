func addBinary(a string, b string) string {
    result := ""
    carry:=0
    i:= len(a)-1
    j:= len(b)-1

    for i>=0 ||j>=0 || carry==1{
        sum:= carry

        if i>=0{
            num,_ := strconv.Atoi(string(a[i]))
            sum += num
            i--
        }

        if j>=0{
            num,_ := strconv.Atoi(string(b[j]))
            sum += num
            j--
        }

        result= strconv.Itoa(sum%2) + result
        carry= sum/2
    }

    return result 
}