package main

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	minHeap := make(MinHeap, k)
	for i := 0; i < k; i++ {
		minHeap[i] = nums[i]
	}

	heap.Init(&minHeap)
	for i := k; i < len(nums); i++ {
		if minHeap[0] < nums[i] {
			minHeap[0] = nums[i]
			heap.Fix(&minHeap, 0)
		}
	}
	return minHeap[0]
}
