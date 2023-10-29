package main

import (
	intHeap "GoBasics/pkg/heap"
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	minHeap := make(intHeap.MinHeap, k)
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
