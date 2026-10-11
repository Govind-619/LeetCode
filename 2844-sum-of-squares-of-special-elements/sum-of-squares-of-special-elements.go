func sumOfSquares(nums []int) int {
    sum:=0
    n:= len(nums)
    b:=1
    for i:=0; i<n; i++{
        if n%b == 0{
           sum= sum + (nums[i]*nums[i])
        }
        b++
    }
 return sum
}