import (
    "math"
)

func largestTriangleArea(points [][]int) float64 {
    maxS := 0.0
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            for k := j + 1; k < len(points); k++ {
                // corners coordinates
                x1 := float64(points[i][0])
                y1 := float64(points[i][1])
                x2 := float64(points[j][0])
                y2 := float64(points[j][1])
                x3 := float64(points[k][0])
                y3 := float64(points[k][1])

                // sides of the triangle
                a := math.Sqrt(math.Pow((x1 - x2), 2) + math.Pow((y1 - y2), 2))
                b := math.Sqrt(math.Pow((x1 - x3), 2) + math.Pow((y1 - y3), 2))
                c := math.Sqrt(math.Pow((x2 - x3), 2) + math.Pow((y2 - y3), 2))

                // Linear equation y = k * x + n
                k := (y1 - y2) / (x1 - x2)
                n := y1 - k * x1
                // if three point are on the same line, we cannot form a triangle
                if y3 == k * x3 + n {
                    continue
                }

                // semiperimeter
                p := (a + b + c) / 2
                // Heron's formula
                s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
                if s > maxS {
                    maxS = s
                }
            }
        }   
    }
    return maxS
}