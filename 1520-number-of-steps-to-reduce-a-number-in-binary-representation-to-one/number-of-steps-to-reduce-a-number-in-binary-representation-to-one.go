func numSteps(s string) int {
	end := strings.LastIndex(s, "1") + 1
	tailZeros := len(s) - end
	if end == 1 {
		return tailZeros
	}
	return end*2 - strings.Count(s[:end], "1") + tailZeros + 1
}