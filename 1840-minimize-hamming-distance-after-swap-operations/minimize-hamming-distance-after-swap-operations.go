func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	dsu := NewDSU(n)

	// Build connected components
	for _, swap := range allowedSwaps {
		dsu.Union(swap[0], swap[1])
	}
	// Group Indices by root
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := dsu.Find(i)
		groups[root] = append(groups[root], i)
	}

	// count matches
	distance := 0
	for _, indices := range groups {
		freq := make(map[int]int)
		for _, i := range indices {
			freq[source[i]]++
		}

		// match with target
		for _, i := range indices {
			if freq[target[i]] > 0 {
				freq[target[i]]--
			} else {
				distance++
			}
		}
	}
	return distance
}

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{
		parent: parent,
	}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	px := d.Find(x)
	py := d.Find(y)
	if px != py {
		d.parent[px] = py
	}
}