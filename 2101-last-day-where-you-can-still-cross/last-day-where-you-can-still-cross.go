func latestDayToCross(row int, col int, cells [][]int) int {
    grid := make([][]int, row, row)
    for i := range grid {
        grid[i] = make([]int, col, col)
    }
    for i, cell := range cells {
        grid[cell[0]-1][cell[1]-1] = i+1
    }    
    // fl := isPossible(grid, 1, row, col)
    l, r := 0, row*col
    for l < r {
        m := (l+r)/2
        if isPossible(grid, m, row, col) {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    if isPossible(grid, l, row, col) {
        return l
    }
    return l-1
}

func isPossible(grid [][]int, day int, row, col int) bool {
    var queue [][]int
    isVisited := make(map[int]map[int]bool)
    for j := range grid[0] {
        queue = appendIfPossible(queue, 0, j, grid, day, isVisited)
    }
    for len(queue) > 0 {
        var newQueue [][]int
        for _, item := range queue {
            if item[0]+1 == row {
                return true
            }
            newQueue = appendIfPossible(newQueue, item[0]+1, item[1], grid, day, isVisited)
            newQueue = appendIfPossible(newQueue, item[0]-1, item[1], grid, day, isVisited)
            newQueue = appendIfPossible(newQueue, item[0], item[1]+1, grid, day, isVisited)
            newQueue = appendIfPossible(newQueue, item[0], item[1]-1, grid, day, isVisited)
        }
        queue = newQueue
    }
    return false
}


func appendIfPossible(queue [][]int, i, j int, grid [][]int, day int, isVisited map[int]map[int]bool) [][]int {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
        return queue
    } 
    if grid[i][j] <= day {
        return queue
    }
    if isVisited[i] != nil && isVisited[i][j] {
        return queue
    }
    if isVisited[i] == nil {
        isVisited[i] = make(map[int]bool)
    }
    isVisited[i][j] = true
    return append(queue, []int{i, j})
}