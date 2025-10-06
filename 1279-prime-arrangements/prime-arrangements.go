func numPrimeArrangements(n int) int {
	data := make([]bool, n+1)
	
	var c int
	for i := 2; i <= n; i++ {
		if !data[i] {
			c++
			for j := i; j+i <= n; {
				data[j+i] = true
				j += i
			}
		}
	}

	return factorial(c) * factorial(n-c) % (1e9 + 7)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1) % (1e9 + 7)
}