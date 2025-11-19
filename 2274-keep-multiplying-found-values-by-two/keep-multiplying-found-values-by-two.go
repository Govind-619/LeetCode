func findFinalValue(nums []int, original int) int {
  for _,v:= range nums{
    if v == original{
       return findFinalValue(nums,original*2)
    }
  } 
  return original 
}