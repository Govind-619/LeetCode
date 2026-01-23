// Complexity:
// Time O(NLogN) and Space O(N), where N is the length of nums.
func minimumPairRemoval(nums []int) int {
	n := len(nums)

	violations := countViolations(nums)
	if violations == 0 {
		return 0
	}

	h := newNodeMinHeap(nums)
	for violations > 0 {
		curr := heap.Pop(h).(*listNode)
		prev := curr.prev
		next := curr.next
		curr.prev = nil
		curr.next = nil
		prev.next = next
		next.prev = prev

		if prev.arrValue > curr.arrValue {
			violations--
		}
		if curr.arrValue > next.arrValue {
			violations--
		}
		if next.arrValue > next.next.arrValue {
			violations--
		}

		if prev.arrIndex != -1 {
			prev.pairSum += next.arrValue
			heap.Fix(h, prev.heapIndex)
		}

		next.arrValue += curr.arrValue
		if next.arrIndex != n-1 {
			next.pairSum += curr.arrValue
			heap.Fix(h, next.heapIndex)
		}

		if prev.arrValue > next.arrValue {
			violations++
		}
		if next.arrValue > next.next.arrValue {
			violations++
		}
	}

	return n - 1 - h.Len()
}

func countViolations(nums []int) int {
	violations := 0
	for i := range len(nums) - 1 {
		if nums[i] > nums[i+1] {
			violations++
		}
	}

	return violations
}

type listNode struct {
	arrIndex  int
	arrValue  int64
	pairSum   int64
	prev      *listNode
	next      *listNode
	heapIndex int
}

type nodeMinHeap []*listNode

func newNodeMinHeap(nums []int) *nodeMinHeap {
	n := len(nums)
	h := make(nodeMinHeap, n)

	var prev *listNode
	prev = &listNode{
		arrIndex:  -1,
		arrValue:  math.MinInt64,
		pairSum:   0,
		prev:      prev,
		next:      nil,
		heapIndex: -1,
	}

	for i, num := range nums {
		h[i] = &listNode{
			arrIndex:  i,
			arrValue:  int64(num),
			pairSum:   int64(num),
			prev:      prev,
			next:      nil,
			heapIndex: i,
		}
		prev.pairSum += int64(num)
		prev.next = h[i]
		prev = h[i]
	}

	h[n-1].next = h[n-1]
	h = h[:n-1]

	heap.Init(&h)
	return &h
}

func (h nodeMinHeap) Len() int {
	return len(h)
}

func (h nodeMinHeap) Less(i, j int) bool {
	if h[i].pairSum == h[j].pairSum {
		return h[i].arrIndex < h[j].arrIndex
	}

	return h[i].pairSum < h[j].pairSum
}

func (h nodeMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex = i
	h[j].heapIndex = j
}

func (h *nodeMinHeap) Push(x any) {
	n := len(*h)
	item := x.(*listNode)
	item.heapIndex = n
	*h = append(*h, item)
}

func (h *nodeMinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	item.heapIndex = -1
	old[n-1] = nil
	*h = old[:n-1]
	return item
}