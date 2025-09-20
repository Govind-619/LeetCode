package main

import (
    "fmt"
    "sort"
)

type Packet struct {
    source      int
    destination int
    timestamp   int
}

type Router struct {
    memoryLimit int
    size        int
    queue       []*Packet // acts as FIFO queue
    front       int

    seen map[string]bool

    destTimestamps   map[int][]int
    destForwardCount map[int]int
}

func Constructor(memoryLimit int) Router {
    return Router{
        memoryLimit:      memoryLimit,
        size:             0,
        queue:            make([]*Packet, 0),
        front:            0,
        seen:             make(map[string]bool),
        destTimestamps:   make(map[int][]int),
        destForwardCount: make(map[int]int),
    }
}

func makeKey(source, destination, timestamp int) string {
    return fmt.Sprintf("%d_%d_%d", source, destination, timestamp)
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
    key := makeKey(source, destination, timestamp)
    if this.seen[key] {
        return false
    }
    if this.size == this.memoryLimit {
        this.forwardOne()
    }
    pkt := &Packet{source, destination, timestamp}
    this.queue = append(this.queue, pkt)
    this.size++
    this.seen[key] = true
    this.destTimestamps[destination] = append(this.destTimestamps[destination], timestamp)
    return true
}

func (this *Router) forwardOne() *Packet {
    if this.size == 0 {
        return nil
    }
    pkt := this.queue[this.front]
    this.front++
    this.size--
    key := makeKey(pkt.source, pkt.destination, pkt.timestamp)
    delete(this.seen, key)
    this.destForwardCount[pkt.destination]++
    return pkt
}

func (this *Router) ForwardPacket() []int {
    pkt := this.forwardOne()
    if pkt == nil {
        return []int{}
    }
    return []int{pkt.source, pkt.destination, pkt.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    arr, ok := this.destTimestamps[destination]
    if !ok {
        return 0
    }
    startIdx := this.destForwardCount[destination]
    if startIdx >= len(arr) {
        return 0
    }

    // binary search left boundary (>= startTime)
    lo := sort.Search(len(arr)-startIdx, func(i int) bool {
        return arr[startIdx+i] >= startTime
    })
    // binary search right boundary (> endTime)
    hi := sort.Search(len(arr)-startIdx, func(i int) bool {
        return arr[startIdx+i] > endTime
    })
    return hi - lo
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
