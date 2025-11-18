// Complexity:
// Time O(N) and Space O(1) where N is the length of bits.
func isOneBitCharacter(bits []int) bool {
	i := len(bits) - 2
	for i >= 0 && bits[i] == 1 {
		i--
	}
	return (len(bits)-i)&1 == 0
}