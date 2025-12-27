func heapPush[E cmp.Ordered](heap []E, item E) []E {
    heap = append(heap, item)
    siftDown(heap, 0, len(heap)-1)
    return heap
}

func heapPop[E cmp.Ordered](heap []E) (E, []E) {
    var lastelt E
    lastelt, heap = heap[len(heap)-1], heap[:len(heap)-1]
    if len(heap) > 0 {
        returnitem := heap[0]
        heap[0] = lastelt
        siftUp(heap, 0)
        return returnitem, heap
    }
    return lastelt, heap
}

func heapReplace[E cmp.Ordered](heap []E, item E) E {
    returnitem := heap[0]
    heap[0] = item
    siftUp(heap, 0)
    return returnitem
}

func siftDown[E cmp.Ordered](heap []E, startpos, pos int) {
    newitem := heap[pos]
    for pos > startpos {
        parentpos := (pos - 1) >> 1
        parent := heap[parentpos]
        if newitem < parent {
            heap[pos] = parent
            pos = parentpos
            continue
        }
        break
    }
    heap[pos] = newitem
}

func siftUp[E cmp.Ordered](heap []E, pos int) {
    endpos := len(heap)
    startpos := pos
    newitem := heap[pos]
    childpos := 2*pos + 1
    for childpos < endpos {
        rightpos := childpos + 1
        if rightpos < endpos && !(heap[childpos] < heap[rightpos]) {
            childpos = rightpos
        }
        heap[pos] = heap[childpos]
        pos = childpos
        childpos = 2*pos + 1
    }
    heap[pos] = newitem
    siftDown(heap, startpos, pos)
}

func room(x uint64) uint8 {
    return uint8(x)
}

func endTime(x uint64) uint64 {
    return x >> 8
}

func item(e uint64, r uint8) uint64 {
    return (e << 8) | uint64(r)
}

func mostBooked(n int, meetings [][]int) int {
    slices.SortFunc(meetings, func (a, b []int) int { return a[0] - b[0] })
    S := make([]int, n)
    H := make([]uint64, 0, n)
    R := make([]uint8, n)
    for i := range R {
        R[i] = uint8(i)
    }
    for _, m := range meetings {
        starti, endi := m[0], m[1]
        for len(H) > 0 && endTime(H[0]) <= uint64(starti) {
            var x uint64
            x, H = heapPop(H)
            R = heapPush(R, room(x))
        }
        if len(R) > 0 {
            var r uint8
            r, R = heapPop(R)
            H = heapPush(H, item(uint64(endi), r))
            S[r]++
        } else {
            e, r := endTime(H[0]), room(H[0])
            heapReplace(H, item(e + uint64(endi - starti), r))
            S[r]++
        }
    }
    m, r := -1, 0
    for i, s := range S {
        if m < s {
            m = s
            r = i
        }
    }
    return r
}
