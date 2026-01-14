type Event struct {
	y     float64
	x1    float64
	x2    float64
	typ   int // +1 start, -1 end
}

type Segment struct {
	yStart     float64
	yEnd       float64
	xCovered   float64
	areaBefore float64
}

type SegmentTree struct {
	cnt   []int
	len   []float64
	xs    []float64
	n     int
}

func NewSegmentTree(xs []float64) *SegmentTree {
	n := len(xs) - 1
	return &SegmentTree{
		cnt: make([]int, 4*n),
		len: make([]float64, 4*n),
		xs:  xs,
		n:   n,
	}
}

func (st *SegmentTree) query() float64 {
	return st.len[1]
}

func (st *SegmentTree) update(node, l, r, ql, qr, val int) {
	if qr < l || r < ql {
		return
	}

	if ql <= l && r <= qr {
		st.cnt[node] += val
	} else {
		mid := (l + r) / 2
		st.update(node*2, l, mid, ql, qr, val)
		st.update(node*2+1, mid+1, r, ql, qr, val)
	}

	if st.cnt[node] > 0 {
		st.len[node] = st.xs[r+1] - st.xs[l]
	} else if l == r {
		st.len[node] = 0
	} else {
		st.len[node] = st.len[node*2] + st.len[node*2+1]
	}
}

func unique(arr []float64) []float64 {
	m := make(map[float64]bool)
	var res []float64
	for _, v := range arr {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func indexOf(arr []float64, target float64) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := (l + r) / 2
		if arr[m] == target {
			return m
		} else if arr[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func separateSquares(squares [][]int) float64 {
	var events []Event
	var xs []float64

	for _, s := range squares {
		x := float64(s[0])
		y := float64(s[1])
		l := float64(s[2])

		events = append(events, Event{y, x, x + l, 1})
		events = append(events, Event{y + l, x, x + l, -1})

		xs = append(xs, x, x+l)
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].y == events[j].y {
			return events[i].typ > events[j].typ
		}
		return events[i].y < events[j].y
	})

	sort.Float64s(xs)
	xs = unique(xs)

	st := NewSegmentTree(xs)

	var segments []Segment
	var totalArea float64

	prevY := events[0].y

	for _, e := range events {
		curY := e.y
		if curY > prevY {
			coveredX := st.query()
			segments = append(segments, Segment{
				yStart:     prevY,
				yEnd:       curY,
				xCovered:   coveredX,
				areaBefore: totalArea,
			})
			totalArea += coveredX * (curY - prevY)
			prevY = curY
		}

		l := indexOf(xs, e.x1)
		r := indexOf(xs, e.x2) - 1
		if l <= r {
			st.update(1, 0, st.n-1, l, r, e.typ)
		}
	}

	target := totalArea / 2

	for _, seg := range segments {
		if seg.xCovered == 0 {
			continue
		}
		area := seg.xCovered * (seg.yEnd - seg.yStart)
		if seg.areaBefore+area >= target {
			return seg.yStart + (target-seg.areaBefore)/seg.xCovered
		}
	}

	return 0
}
