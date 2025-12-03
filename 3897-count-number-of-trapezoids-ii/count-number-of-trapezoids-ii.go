// Helper function to compute GCD
func gcd(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return gcd(b, a%b)
}

// Normalize dy/dx to reduced form with consistent sign
func normalize(a, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	if a == 0 {
		return 0, 1
	}
	g := gcd(a, b)
	a, b = a/g, b/g
	if b < 0 {
		a, b = -a, -b
	}
	return a, b
}

type slopeKey struct {
	dy, dx int
}
type lineKey struct {
	dy, dx, cdy int // cdy = dx*y1 - dy*x1
}
type midKey struct {
	sx, sy int
}
type fullKey struct {
	sx, sy, dy, dx, cdy int
}

func countTrapezoids(points [][]int) int {
	slopes := map[slopeKey]int{}
	lines := map[lineKey]int{}
	mids := map[midKey]int{}
	full := map[fullKey]int{}

	for i, a := range points {
		for j := 0; j < i; j++ {
			b := points[j]
			x1, y1 := a[0], a[1]
			x2, y2 := b[0], b[1]

			dy, dx := y2-y1, x2-x1
			ny, nx := normalize(dy, dx)

			cdy := nx*y1 - ny*x1
			midx, midy := x1+x2, y1+y2

			slopes[slopeKey{ny, nx}]++
			lines[lineKey{ny, nx, cdy}]++
			mids[midKey{midx, midy}]++
			full[fullKey{midx, midy, ny, nx, cdy}]++
		}
	}

	ans := 0
	for _, v := range slopes {
		ans += v * (v - 1) / 2
	}
	for _, v := range lines {
		ans -= v * (v - 1) / 2
	}
	for _, v := range mids {
		ans -= v * (v - 1) / 2
	}
	for _, v := range full {
		ans += v * (v - 1) / 2
	}
	return ans
}