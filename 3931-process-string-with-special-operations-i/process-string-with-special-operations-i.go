package main

func processStr(s string) string {
	// Pre-allocate a reasonable capacity to avoid early re-allocations
	res := make([]byte, 0, len(s)*2)

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch ch {
		case '*':
			if len(res) > 0 {
				res = res[:len(res)-1] // O(1) pop
			}
			
		case '#':
			// Go's built-in append uses highly optimized memory copying under the hood.
			// This safely handles exponential growth without any bounds panics.
			res = append(res, res...)
			
		case '%':
			// Physical reverse in a flat slice.
			// Because memory is contiguous, the CPU cache absolutely crushes this loop.
			for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
				res[l], res[r] = res[r], res[l]
			}
			
		default:
			res = append(res, ch) // Amortized O(1) append
		}
	}

	return string(res)
}