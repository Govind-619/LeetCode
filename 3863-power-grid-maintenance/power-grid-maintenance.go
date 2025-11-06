import (
    "container/heap"
)

type DSU struct {
    parent []int
}

func NewDSU(n int) *DSU {
    p := make([]int, n+1)
    for i := 1; i <= n; i++ {
        p[i] = i
    }
    return &DSU{parent: p}
}

func (d *DSU) Find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.Find(d.parent[x])
    }
    return d.parent[x]
}

func (d *DSU) Union(x, y int) {
    rx, ry := d.Find(x), d.Find(y)
    if rx != ry {
        d.parent[ry] = rx
    }
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
    dsu := NewDSU(c)
    for _, e := range connections {
        dsu.Union(e[0], e[1])
    }

    online := make([]bool, c+1)
    for i := 1; i <= c; i++ {
        online[i] = true
    }

    // To guarantee heap grouping is by unique root after path compression
    for i := 1; i <= c; i++ {
        dsu.Find(i)
    }

    // Build mapping from root → member stations
    comp := make(map[int]*MinHeap)
    for i := 1; i <= c; i++ {
        root := dsu.parent[i] // path-compressed root
        if comp[root] == nil {
            comp[root] = &MinHeap{}
            heap.Init(comp[root])
        }
        heap.Push(comp[root], i)
    }

    var res []int

    for _, q := range queries {
        t, x := q[0], q[1]
        if t == 2 {
            online[x] = false
        } else if t == 1 {
            root := dsu.Find(x)
            h := comp[root]
            // Clean lazy-deleted values
            for h.Len() > 0 && !online[(*h)[0]] {
                heap.Pop(h)
            }
            if online[x] {
    res = append(res, x)
} else if h.Len() == 0 {
    res = append(res, -1)
} else {
    res = append(res, (*h)[0])
}
        }
    }
    return res
}
