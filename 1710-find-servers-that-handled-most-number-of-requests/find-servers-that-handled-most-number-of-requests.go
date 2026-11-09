type entry struct{ key, server int }  // key: the next turn in free, the time it frees up in busy

type minHeap []entry

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }
func (h minHeap) Less(a, b int) bool { return h[a].key < h[b].key }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(entry)) }
func (h *minHeap) Pop() any {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}

func busiestServers(k int, arrival []int, load []int) []int {
    free := make(minHeap, k)
    for x := range free {
        free[x] = entry{x, x}  // at request 0, server x's next turn is x
    }
    heap.Init(&free)

    handled, busy := make([]int, k), minHeap{}
    for i, t := range arrival {
        for len(busy) > 0 && busy[0].key <= t {  // done by the arrival
            x := heap.Pop(&busy).(entry).server
            heap.Push(&free, entry{i + (x-i%k+k)%k, x})  // its first turn from i on
        }
        if len(free) == 0 {  // none free: the request is dropped
            continue
        }
        x := heap.Pop(&free).(entry).server
        heap.Push(&busy, entry{t + load[i], x})
        handled[x]++
    }

    most := slices.Max(handled)
    ans := []int{}
    for x, h := range handled {
        if h == most {
            ans = append(ans, x)
        }
    }
    return ans
}