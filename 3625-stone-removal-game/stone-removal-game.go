func canAliceWin(n int) bool {
    missing := 55 - n  // stones short of playing all ten moves
    unplayed := int(math.Round(math.Sqrt(float64(2 * missing))))  // smallest j with 1 + 2 + ... + j >= missing
    return unplayed%2 == 1  // 10 - unplayed moves were played: odd means Alice moved last
}