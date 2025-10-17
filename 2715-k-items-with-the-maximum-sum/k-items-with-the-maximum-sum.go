func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
    if k<=numOnes{
        return k
    }else if k<= numOnes + numZeros{
        return numOnes
    }else{
         val:= numOnes
         k= k-numOnes- numZeros
         val += k*-1
         return val
    }
}