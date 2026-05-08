const MX = 1000001
var isPrime [MX]bool

// Precompute primes once (static initialization)
func init() {
    for i := 2; i < MX; i++ {
        isPrime[i] = true
    }
    for i := 2; i*i < MX; i++ {
        if isPrime[i] {
            for j := i * i; j < MX; j += i{
                isPrime[j] = false
            }
        }
    }
}

func minJumps(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 0
    }

    // Map each prime 'p' to all indices 'j' where nums[j] is divisible by 'p'
    // This allows us to find jump targets quickly
    primeToIndices := make(map[int][]int)
    for i, x := range nums {
        temp := x
        for d := 2; d*d <= temp; d++ {
            if temp%d == 0 {
                primeToIndices[d] = append(primeToIndices[d], i)
                for temp%d == 0 {
                    temp /= d
                }
            }
        }
        if temp > 1 {
            primeToIndices[temp] = append(primeToIndices[temp], i)
        }
    }

    queue := []int{0}
    visitedIdx := make([]bool, n)
    visitedIdx[0] = true
    steps := 0

    for len(queue) > 0 {
        size := len(queue)
        for k := 0; k < size; k++ {
            curr := queue[0]
            queue = queue[1:]

            if curr == n-1 {
                return steps
            }

            // 1. Adjacent Steps (i+1, i-1)
            for _, next := range []int{curr - 1, curr + 1} {
                if next >= 0 && next < n && !visitedIdx[next] {
                    visitedIdx[next] = true
                    queue = append(queue, next)
                }
            }

            // 2. Prime Teleportation (ONLY if nums[curr] is prime)
            p := nums[curr]
            if p < MX && isPrime[p] {
                if indices, ok := primeToIndices[p]; ok {
                    for _, idx := range indices {
                        if !visitedIdx[idx] {
                            visitedIdx[idx] = true
                            queue = append(queue, idx)
                        }
                    }
                    // Crucial: Clear the bucket after visiting to ensure O(N)
                    delete(primeToIndices, p)
                }
            }
        }
        steps++
    }

    return -1
}