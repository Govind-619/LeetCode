func maxBottlesDrunk(numBottles int, numExchange int) int {
    // You can initially drink all bottles you have
    ans := numBottles

    // While you have enough empty bottles to trade
    for numBottles >= numExchange {
        // Perform one exchange:
        // You spend `numExchange` empties to get 1 full bottle,
        // then you drink it => that gives 1 new empty bottle
        numBottles = numBottles - numExchange + 1

        // Since numExchange increases by 1 after each trade
        numExchange++

        // You drank one more bottle
        ans++
    }

    return ans
}
