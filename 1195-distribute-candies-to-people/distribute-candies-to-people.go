func distributeCandies(candies int, num_people int) []int {
    res := make([]int, num_people)
    n := 1
    for candies > 0 {
        for i := 0; i < num_people && candies > 0; i++ {
            if candies < n {
                res[i] += candies
                candies = 0
                break
            }
            res[i] += n
            candies -= n
            n++
        }
    }
    return res
}
