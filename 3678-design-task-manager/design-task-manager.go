package main

import (
    "container/heap"
)

// Task is an element in the heap
type Task struct {
    taskId   int
    priority int
}

// MaxHeap implements heap.Interface (max-heap by priority, then taskId)
type MaxHeap []Task

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
    if h[i].priority == h[j].priority {
        return h[i].taskId > h[j].taskId
    }
    return h[i].priority > h[j].priority
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(Task))
}
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    t := old[n-1]
    *h = old[:n-1]
    return t
}

type TaskInfo struct {
    userId   int
    priority int
}

type TaskManager struct {
    taskMap map[int]TaskInfo  // taskId -> (userId, priority)
    heap    *MaxHeap
}

func Constructor(tasks [][]int) TaskManager {
    h := &MaxHeap{}
    heap.Init(h)
    tm := TaskManager{
        taskMap: make(map[int]TaskInfo),
        heap:    h,
    }
    for _, t := range tasks {
        userId := t[0]
        taskId := t[1]
        pr := t[2]
        tm.taskMap[taskId] = TaskInfo{userId, pr}
        heap.Push(h, Task{taskId, pr})
    }
    return tm
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
    tm.taskMap[taskId] = TaskInfo{userId, priority}
    heap.Push(tm.heap, Task{taskId, priority})
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
    if info, ok := tm.taskMap[taskId]; ok {
        // update map
        tm.taskMap[taskId] = TaskInfo{info.userId, newPriority}
        // push new version into heap
        heap.Push(tm.heap, Task{taskId, newPriority})
    }
}

func (tm *TaskManager) Rmv(taskId int) {
    // Simply delete it from map; any entries in heap become stale
    delete(tm.taskMap, taskId)
}

func (tm *TaskManager) ExecTop() int {
    // Pop until we find a valid top
    for tm.heap.Len() > 0 {
        top := (*tm.heap)[0]
        // Check if it's still in map and priority matches
        if info, ok := tm.taskMap[top.taskId]; ok && info.priority == top.priority {
            // valid
            userId := info.userId
            // remove from map
            delete(tm.taskMap, top.taskId)
            // remove from heap
            heap.Pop(tm.heap)
            return userId
        }
        // else: stale entry, pop and continue
        heap.Pop(tm.heap)
    }
    return -1
}
