package main

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
    left, right := int64(0), int64(1e18) // right is initialized to a large value
    res := right

    for left <= right {
        mid := left + (right-left)/2 // Calculate the mid value

        if canComplete(mid, mountainHeight, workerTimes) {
            res = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return res
}

func canComplete(timeLimit int64, mountainHeight int, workerTimes []int) bool {
    totalReduction := int64(0)

    for _, t := range workerTimes {
        low, high := int64(0), int64(mountainHeight)

        for low <= high {
            mid := low + (high-low)/2
            timeRequired := int64(t) * (mid * (mid + 1)) / 2

            if timeRequired <= timeLimit {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }

        totalReduction += high
        if totalReduction >= int64(mountainHeight) {
            return true
        }
    }
    return totalReduction >= int64(mountainHeight)
}