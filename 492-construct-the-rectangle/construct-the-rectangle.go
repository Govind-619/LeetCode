func constructRectangle(area int) []int {
    w := int(math.Sqrt(float64(area)))
    for w > 0 {
        if area%w == 0 {
            l := area / w
            return []int{l, w} // L >= W
        }
        w--
    }
    return []int{area, 1} // fallback, e.g., area is prime
}