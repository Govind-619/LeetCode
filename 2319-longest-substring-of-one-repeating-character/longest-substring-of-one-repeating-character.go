func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	type Node struct {
		leftChar  byte
		rightChar byte
		len       int
		prefix    int
		suffix    int
		best      int
	}

	n := len(s)

	tree := make([]Node, 4*n)

	merge := func(left, right Node) Node {
		var res Node

		res.leftChar = left.leftChar
		res.rightChar = right.rightChar
		res.len = left.len + right.len

		res.best = left.best
		if right.best > res.best {
			res.best = right.best
		}

		res.prefix = left.prefix

		if left.prefix == left.len &&
			left.rightChar == right.leftChar {
			res.prefix = left.len + right.prefix
		}

		res.suffix = right.suffix

		if right.suffix == right.len &&
			left.rightChar == right.leftChar {
			res.suffix = left.suffix + right.len
		}

		if left.rightChar == right.leftChar {
			cross := left.suffix + right.prefix
			if cross > res.best {
				res.best = cross
			}
		}

		return res
	}

	var build func(int, int, int)

	build = func(node, l, r int) {
		if l == r {
			tree[node] = Node{
				leftChar:  s[l],
				rightChar: s[l],
				len:       1,
				prefix:    1,
				suffix:    1,
				best:      1,
			}
			return
		}

		mid := (l + r) / 2

		build(node*2, l, mid)
		build(node*2+1, mid+1, r)

		tree[node] = merge(tree[node*2], tree[node*2+1])
	}

	var update func(int, int, int, int, byte)

	update = func(node, l, r, idx int, c byte) {
		if l == r {
			tree[node] = Node{
				leftChar:  c,
				rightChar: c,
				len:       1,
				prefix:    1,
				suffix:    1,
				best:      1,
			}
			return
		}

		mid := (l + r) / 2

		if idx <= mid {
			update(node*2, l, mid, idx, c)
		} else {
			update(node*2+1, mid+1, r, idx, c)
		}

		tree[node] = merge(tree[node*2], tree[node*2+1])
	}

	build(1, 0, n-1)

	answer := make([]int, len(queryCharacters))

	for i := 0; i < len(queryCharacters); i++ {
		update(
			1,
			0,
			n-1,
			queryIndices[i],
			queryCharacters[i],
		)

		answer[i] = tree[1].best
	}

	return answer
}