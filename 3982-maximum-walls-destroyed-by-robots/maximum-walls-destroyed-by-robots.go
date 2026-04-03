package main

import (
	"sort"
)

func maxWalls(robots []int, distance []int, walls []int) int {
	n := len(robots)

	left := make([]int, n)
	right := make([]int, n)
	num := make([]int, n)

	// map robot -> distance
	robotsToDistance := make(map[int]int)
	for i := 0; i < n; i++ {
		robotsToDistance[robots[i]] = distance[i]
	}

	sort.Ints(robots)
	sort.Ints(walls)

	for i := 0; i < n; i++ {
		// upper_bound(walls, robots[i])
		pos1 := sort.Search(len(walls), func(j int) bool {
			return walls[j] > robots[i]
		})

		var leftPos int
		if i >= 1 {
			leftBound := max(robots[i]-robotsToDistance[robots[i]], robots[i-1]+1)
			leftPos = sort.Search(len(walls), func(j int) bool {
				return walls[j] >= leftBound
			})
		} else {
			leftBound := robots[i] - robotsToDistance[robots[i]]
			leftPos = sort.Search(len(walls), func(j int) bool {
				return walls[j] >= leftBound
			})
		}
		left[i] = pos1 - leftPos

		var rightPos int
		if i < n-1 {
			rightBound := min(robots[i]+robotsToDistance[robots[i]], robots[i+1]-1)
			rightPos = sort.Search(len(walls), func(j int) bool {
				return walls[j] > rightBound
			})
		} else {
			rightBound := robots[i] + robotsToDistance[robots[i]]
			rightPos = sort.Search(len(walls), func(j int) bool {
				return walls[j] > rightBound
			})
		}

		// lower_bound(walls, robots[i])
		pos2 := sort.Search(len(walls), func(j int) bool {
			return walls[j] >= robots[i]
		})

		right[i] = rightPos - pos2

		if i == 0 {
			continue
		}

		// lower_bound(walls, robots[i-1])
		pos3 := sort.Search(len(walls), func(j int) bool {
			return walls[j] >= robots[i-1]
		})

		num[i] = pos1 - pos3
	}

	// DP
	subLeft := left[0]
	subRight := right[0]

	for i := 1; i < n; i++ {
		currentLeft := max(
			subLeft+left[i],
			subRight-right[i-1]+min(left[i]+right[i-1], num[i]),
		)

		currentRight := max(
			subLeft+right[i],
			subRight+right[i],
		)

		subLeft = currentLeft
		subRight = currentRight
	}

	return max(subLeft, subRight)
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