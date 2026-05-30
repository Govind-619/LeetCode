type Fenwick struct {
	n   int
	bit []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		n:   n,
		bit: make([]int, n+1),
	}
}

func (f *Fenwick) Add(idx, val int) {
	for idx <= f.n {
		f.bit[idx] += val
		idx += idx & -idx
	}
}

func (f *Fenwick) Sum(idx int) int {
	res := 0
	for idx > 0 {
		res += f.bit[idx]
		idx -= idx & -idx
	}
	return res
}

func (f *Fenwick) Kth(k int) int {
	idx := 0
	step := 1
	for (step << 1) <= f.n {
		step <<= 1
	}
	for d := step; d > 0; d >>= 1 {
		next := idx + d
		if next <= f.n && f.bit[next] < k {
			idx = next
			k -= f.bit[next]
		}
	}
	return idx + 1
}

type SegTree struct {
	n    int
	tree []int
}

func NewSegTree(n int) *SegTree {
	size := 4
	if 4*n > size {
		size = 4 * n
	}
	return &SegTree{
		n:    n,
		tree: make([]int, size),
	}
}

func (s *SegTree) update(node, l, r, pos, val int) {
	if l == r {
		s.tree[node] = val
		return
	}
	mid := (l + r) >> 1
	if pos <= mid {
		s.update(node<<1, l, mid, pos, val)
	} else {
		s.update(node<<1|1, mid+1, r, pos, val)
	}
	left, right := s.tree[node<<1], s.tree[node<<1|1]
	if left > right {
		s.tree[node] = left
	} else {
		s.tree[node] = right
	}
}

func (s *SegTree) query(node, l, r, ql, qr int) int {
	if ql > r || qr < l {
		return 0
	}
	if ql <= l && r <= qr {
		return s.tree[node]
	}
	mid := (l + r) >> 1
	a := s.query(node<<1, l, mid, ql, qr)
	b := s.query(node<<1|1, mid+1, r, ql, qr)
	if a > b {
		return a
	}
	return b
}

func getResults(queries [][]int) []bool {
	mx := 0
	for _, q := range queries {
		if q[1] > mx {
			mx = q[1]
		}
	}

	fw := NewFenwick(mx + 2)
	st := NewSegTree(mx + 1)

	fw.Add(1, 1)

	ans := make([]bool, 0, len(queries))

	for _, q := range queries {
		t := q[0]
		x := q[1]

		if t == 1 {
			leftCount := fw.Sum(x)
			leftPos := fw.Kth(leftCount) - 1

			occupiedUpToX := fw.Sum(x + 1)
			totalOccupied := fw.Sum(mx + 2)
			rightPos := -1
			if occupiedUpToX < totalOccupied {
				rightPos = fw.Kth(occupiedUpToX + 1) - 1
			}

			st.update(1, 0, mx, x, x-leftPos)

			if rightPos != -1 {
				st.update(1, 0, mx, rightPos, rightPos-x)
			}

			fw.Add(x+1, 1)
		} else {
			sz := q[2]

			leftCount := fw.Sum(x)
			leftPos := fw.Kth(leftCount) - 1

			bestPrefix := st.query(1, 0, mx, 0, x)

			ans = append(ans, x-leftPos >= sz || bestPrefix >= sz)
		}
	}

	return ans
}