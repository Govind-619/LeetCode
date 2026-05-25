func canReach(s string, minJump int, maxJump int) bool {

    n := len(s)

    queue := []int{0}

    visited := make([]bool, n)
    visited[0] = true

    front := 0

    far := 0

    for front < len(queue) {

        i := queue[front]
        front++

        if i == n-1 {
            return true
        }

        start := max(i+minJump, far+1)
        end := min(i+maxJump, n-1)

        for j := start; j <= end; j++ {

            if s[j] == '0' && !visited[j] {
                visited[j] = true
                queue = append(queue, j)
            }
        }

        if end > far {
            far = end
        }
    }

    return false
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}