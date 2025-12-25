func finalPrices(prices []int) []int {
    res:= make([]int, len(prices))
    copy(res,prices)

    for i:= 0; i< len(prices); i++{
        for j:=i+1; j<len(prices); j++{
            if prices[j] <= prices[i]{
                res[i]=prices[i]- prices[j]
                break
            }
        } 
    }
    return res
}