func robotSim(commands []int, obstacles [][]int) int {
    obstacleSet := make(map[string]bool)

    for _, obs := range obstacles {
        key := fmt.Sprintf("%d,%d", obs[0], obs[1])
        obstacleSet[key] = true
    }

    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}

    dir := 0
    x, y := 0, 0
    maxDistance := 0

    for _, command := range commands {
        if command == -1 {
            dir = (dir + 1) % 4
        } else if command == -2 {
            dir = (dir + 3) % 4
        } else {
            for step := 0; step < command; step++ {
                nextX := x + dx[dir]
                nextY := y + dy[dir]

                key := fmt.Sprintf("%d,%d", nextX, nextY)

                if obstacleSet[key] {
                    break
                }

                x = nextX
                y = nextY

                distance := x*x + y*y
                if distance > maxDistance {
                    maxDistance = distance
                }
            }
        }
    }

    return maxDistance
}