type Heap [][]int  // items, the cheapest on top

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(item any)     { *h = append(*h, item.([]int)) }
func (h *Heap) Pop() any {
    last := len(*h) - 1
    item := (*h)[last]
    *h = (*h)[:last]
    return item
}

func maximumSaleItems(items [][]int, budget int) int {
    n, cheapest := len(items), items[0][1]
    have := make([]int, n+1)  // have[f]: items with factor f
    for _, item := range items {
        have[item[0]]++
        cheapest = min(cheapest, item[1])
    }
    deals := slices.Repeat([]int{-1}, n+1)  // counted when an item is reached

    worth := Heap{}
    for _, item := range items {
        if item[1] < 2*cheapest {
            worth = append(worth, item)
        }
    }
    heap.Init(&worth)  // built in O(n)
    copies := 0
    for len(worth) > 0 && worth[0][1] <= budget {  // the cheapest deal left is affordable
        item := heap.Pop(&worth).([]int)
        factor, price := item[0], item[1]
        if deals[factor] < 0 {  // from -1: the item itself is not a deal
            for multiple := factor; multiple <= n; multiple += factor {
                deals[factor] += have[multiple]
            }
        }
        pairs := min(deals[factor], budget/price)
        copies += 2 * pairs
        budget -= pairs * price
    }
    return copies + budget/cheapest
}