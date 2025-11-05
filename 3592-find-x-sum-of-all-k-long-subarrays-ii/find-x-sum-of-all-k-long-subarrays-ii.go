import "container/heap"

// item represents a distinct value present in the current sliding window.
//   - val: the integer value
//   - freq: how many times 'val' appears in the current window
//   - idx: index inside whichever heap the item currently resides in (hot or rest),
//     or -1 if it is not currently in any heap (temporarily removed for update).
//   - inTop: whether the item is currently counted in the 'hot' heap (top-x elements)
//
// Note: we keep pointers to items in a map so we can update/remove them efficiently.
type item struct {
	val   int
	freq  int
	idx   int
	inTop bool
}

// hotHeap is a min-heap ordered by frequency, then by value.
// It holds up to x items which are the current 'top x' elements we are summing.
// We use a min-heap so that the "worst" among the top-x (smallest freq/value) can be popped
// quickly when we need to demote an element from hot to rest.
type hotHeap []*item

func (h hotHeap) Len() int { return len(h) }

// Less defines order: smaller frequency is considered "less" (i.e., worse in hot),
// and if frequencies are equal, smaller value is "less". This makes Pop remove
// the worst element in hot (lowest priority among the top).
func (h hotHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq < h[j].freq
	}
	return h[i].val < h[j].val
}

func (h hotHeap) Swap(i, j int) {
	// swap elements and update their recorded indices
	h[i], h[j] = h[j], h[i]
	h[i].idx, h[j].idx = i, j
}

func (h *hotHeap) Push(x any) {
	// Push uses pointer semantics. When pushing, set its idx to its position.
	it := x.(*item)
	it.idx = len(*h)
	*h = append(*h, it)
}

func (h *hotHeap) Pop() any {
	// Pop removes and returns the last element (heap.Pop will have ensured it's the right one).
	old := *h
	it := old[len(old)-1]
	*h = old[:len(old)-1]
	// mark idx as -1 since it's no longer in the heap
	it.idx = -1
	return it
}

// restHeap is a max-heap ordered by frequency, then by value.
// It contains the remaining items that are not part of the current top-x.
// We want the largest-frequency (or largest-value on ties) item available in rest
// so we can promote it into hot when needed.
type restHeap []*item

func (h restHeap) Len() int { return len(h) }

// For restHeap we invert the comparisons so the heap gives us the item with
// the highest frequency (and highest value on tie) when we Pop from rest.
func (h restHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}
	return h[i].val > h[j].val
}

func (h restHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx, h[j].idx = i, j
}

func (h *restHeap) Push(x any) {
	it := x.(*item)
	it.idx = len(*h)
	*h = append(*h, it)
}

func (h *restHeap) Pop() any {
	old := *h
	it := old[len(old)-1]
	*h = old[:len(old)-1]
	it.idx = -1
	return it
}

// findXSum computes, for each sliding window of size k over nums,
// the sum of value * frequency for the top x distinct values by frequency (and value tie-breaker).
//
// Returns an array of length n-k+1 where each entry is the sum for that window.
//
// High level approach:
// - Maintain two heaps:
//   - hot (min-heap): contains up to x items that are currently considered the "top x"
//   - rest (max-heap): all other items
//   - Maintain a map from value to *item so we can update frequencies efficiently.
//   - Maintain a running 'sum' that equals sum_{it in hot} it.val * it.freq.
//   - For insertion/removal within the window, we remove the item from whichever heap it is in,
//     adjust its frequency, and then reinsert it into an appropriate heap.
//   - After each insertion/erasure we make sure hot has exactly min(x, number of distinct items) items,
//     promoting from rest when possible or demoting from hot when hot is too large.
//
// Complexity: O(n log m) where m is the number of distinct elements in a window (log factor from heap ops).
func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
	ans := make([]int64, n-k+1)

	// Map from value to its item pointer so we can update frequency in O(1).
	freq := map[int]*item{}

	hot := &hotHeap{}   // min-heap for current top-x items
	rest := &restHeap{} // max-heap for remaining items
	heap.Init(hot)
	heap.Init(rest)

	var sum int64 // current sum of value * frequency for items in hot

	// removeFromCurrent removes 'it' from whichever heap it currently occupies.
	// - If it.inTop is true, it's currently in hot and we must subtract its contribution from sum.
	// - Otherwise, if it.idx >= 0 it resides in rest and we simply remove it there.
	// This function leaves the item with idx = -1 and inTop = false (not in any heap).
	removeFromCurrent := func(it *item) {
		if it.inTop {
			// If in hot, remove its contribution and remove from hot heap.
			sum -= int64(it.val) * int64(it.freq)
			heap.Remove(hot, it.idx)
			it.inTop = false
			it.idx = -1
		} else if it.idx >= 0 {
			// Item is in rest heap: remove it. No change to sum because rest items are not counted.
			heap.Remove(rest, it.idx)
			it.idx = -1
		}
	}

	// promoteIfPossible moves the best items from rest to hot while hot has fewer than x items.
	// Each promoted item's contribution is added to sum and its inTop flag set.
	promoteIfPossible := func() {
		for hot.Len() < x && rest.Len() > 0 {
			best := heap.Pop(rest).(*item) // rest.Pop gives highest freq/value
			best.inTop = true
			sum += int64(best.val) * int64(best.freq)
			heap.Push(hot, best) // push into hot (min-heap)
		}
	}

	// insertVal increments frequency of v and puts/updates the item into appropriate heap(s).
	// Steps:
	// - If item exists, remove it from its current heap so we can update its freq safely.
	// - Increase frequency.
	// - Attempt to put it into hot first (optimistically adding to top-x).
	// - If hot grows beyond x, demote the worst element from hot to rest and adjust sum.
	insertVal := func(v int) {
		it, ok := freq[v]
		if !ok {
			// first time we see this value in the window: create item with idx=-1 (not in any heap yet)
			it = &item{val: v, idx: -1}
			freq[v] = it
		} else {
			// if existing, remove from whichever heap it occupies before changing freq
			removeFromCurrent(it)
		}
		it.freq++ // increment frequency for this insertion

		// Optimistically insert into hot (we prefer new/updated items to be in hot if possible)
		it.inTop = true
		sum += int64(it.val) * int64(it.freq)
		heap.Push(hot, it)

		// If hot exceeds allowed size x, demote the worst from hot to rest.
		// We must also remove its contribution from sum because rest items are not counted.
		if hot.Len() > x {
			worst := heap.Pop(hot).(*item)
			sum -= int64(worst.val) * int64(worst.freq)
			worst.inTop = false
			heap.Push(rest, worst)
		}
	}

	// eraseVal decrements frequency of v (value moving out of sliding window).
	// Steps:
	// - Remove it from whichever heap it occupies.
	// - Decrease frequency; if it becomes zero delete from map.
	// - Otherwise, push it into rest (it cannot be in hot immediately after decrement).
	// - Finally, try to promote items from rest into hot until hot has x items again.
	eraseVal := func(v int) {
		it := freq[v]
		// Remove from current heap so we can safely change its freq
		removeFromCurrent(it)
		it.freq-- // element removed from window, so decrease frequency

		if it.freq == 0 {
			// No more occurrences in window: remove from map completely
			delete(freq, v)
		} else {
			// It still exists in window but with lower frequency: it belongs to rest by default now.
			it.inTop = false
			heap.Push(rest, it)
		}

		// After decreasing frequency, we might have room in hot; promote best from rest.
		promoteIfPossible()
	}

	// Build initial window of size k
	for i := 0; i < k; i++ {
		insertVal(nums[i])
	}
	// record sum for first window
	ans[0] = sum

	// Slide window one step at a time:
	// - erase the outgoing element (i-k)
	// - insert the incoming element (i)
	// - record the current sum
	for i := k; i < n; i++ {
		eraseVal(nums[i-k]) // remove outgoing
		insertVal(nums[i])  // add incoming
		ans[i-k+1] = sum
	}

	return ans
}