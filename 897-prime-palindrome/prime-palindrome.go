func primePalindrome(n int) int {
    if n <= 2 {
        return 2
    }
    if n <= 3 {
        return 3
    }
    if n <= 5 {
        return 5
    }
    if n <= 7 {
        return 7
    }
    if n <= 11 {
        return 11
    }

    for i := 1; ; i++ {
        p := makePalindrome(i)
        if p >= n && isPrime(p) {
            return p
        }
    }
}

func makePalindrome(x int) int {
    res := x
    x /= 10
    for x > 0 {
        res = res*10 + x%10
        x /= 10
    }
    return res
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n%2 == 0 {
        return n == 2
    }
    for i := 3; i*i <= n; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}
