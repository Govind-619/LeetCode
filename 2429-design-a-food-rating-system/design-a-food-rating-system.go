import (
	"container/heap"
	"fmt"
)

type Item struct {
	food   string
	rating int
}

type MaxHeap []Item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].food < h[j].food // lexicographically smaller first
	}
	return h[i].rating > h[j].rating
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRating  map[string]int
	cuisineFoods  map[string]*MaxHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineFoods:  make(map[string]*MaxHeap),
	}

	for i := range foods {
		fr.foodToCuisine[foods[i]] = cuisines[i]
		fr.foodToRating[foods[i]] = ratings[i]
		if fr.cuisineFoods[cuisines[i]] == nil {
			fr.cuisineFoods[cuisines[i]] = &MaxHeap{}
		}
		heap.Push(fr.cuisineFoods[cuisines[i]], Item{foods[i], ratings[i]})
	}

	return fr
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := fr.foodToCuisine[food]
	fr.foodToRating[food] = newRating
	heap.Push(fr.cuisineFoods[cuisine], Item{food, newRating})
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	h := fr.cuisineFoods[cuisine]
	for {
		top := (*h)[0]
		if fr.foodToRating[top.food] == top.rating {
			return top.food
		}
		heap.Pop(h) // discard stale entry
	}
}
