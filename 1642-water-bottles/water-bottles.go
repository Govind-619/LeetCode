func numWaterBottles(numBottles int, numExchange int) int {
    sum:=0
    fix:= numBottles
    for fix>0{
        sum += fix
        fix = numBottles/numExchange
        numBottles= numBottles/numExchange + numBottles%numExchange
    }
    return sum
}