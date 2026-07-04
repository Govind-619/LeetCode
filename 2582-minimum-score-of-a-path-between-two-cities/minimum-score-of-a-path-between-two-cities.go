func minScore(n int, roads [][]int) int {
    type Edge struct {
        city     int
        distance int
    }

    graph := make([][]Edge, n+1)

    for _, road := range roads {
        a := road[0]
        b := road[1]
        distance := road[2]

        graph[a] = append(graph[a], Edge{b, distance})
        graph[b] = append(graph[b], Edge{a, distance})
    }

    visited := make([]bool, n+1)

    queue := []int{1}
    front := 0
    visited[1] = true

    answer := int(^uint(0) >> 1)

    for front < len(queue) {
        city := queue[front]
        front++

        for _, edge := range graph[city] {
            if edge.distance < answer {
                answer = edge.distance
            }

            if !visited[edge.city] {
                visited[edge.city] = true
                queue = append(queue, edge.city)
            }
        }
    }

    return answer
}