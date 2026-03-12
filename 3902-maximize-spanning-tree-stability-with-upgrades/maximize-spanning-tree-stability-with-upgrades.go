package main

import (
	"sort"
)

func find(i int, ds []int) int {
	if ds[i] < 0 {
		return i
	}
	ds[i] = find(ds[i], ds)
	return ds[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min(a, min(b, c))
}

func maxStability(n int, edges [][]int, k int) int {
	suv := make([][3]int, 0)
	ds := make([]int, n)

	for i := 0; i < n; i++ {
		ds[i] = -1
	}

	usedE := 0
	res := 200000
	minSingle := 200000
	minDouble := 100000

	for _, e := range edges {
		if e[3] == 1 {
			a := find(e[0], ds)
			b := find(e[1], ds)

			if a == b {
				return -1
			} else {
				usedE++
				ds[b] = a
				res = min(res, e[2])
			}
		} else {
			suv = append(suv, [3]int{e[2], e[1], e[0]})
		}
	}

	sort.Slice(suv, func(i, j int) bool {
		return suv[i][0] > suv[j][0]
	})

	for _, item := range suv {
		s := item[0]
		u := item[1]
		v := item[2]

		a := find(u, ds)
		b := find(v, ds)

		if a != b {
			ds[b] = a
			usedE++

			if usedE == n-1-k {
				minSingle = s
			}
			minDouble = s
		}
	}

	if usedE == n-1 {
		if k > 0 {
			return min3(res, minSingle, minDouble*2)
		}
		return min3(res, minSingle, minDouble)
	}

	return -1
}
