package algorithms

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	// compute num
	intMap := make(map[int]int)
	for i := range nums {
		k := nums[i]
		_, has := intMap[k]
		if has {
			intMap[k]++
			continue
		}
		intMap[k] = 1
	}

	// push heap
	h := new(intHeap)
	for i := range intMap {
		heap.Push(h, entity{num: intMap[i], value: i})
	}

	// pop topK
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).(entity).value)
	}
	return result
}

type entity struct {
	num   int
	value int
}

type intHeap []entity

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i].num > h[j].num
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(entity))
}

func (h *intHeap) Pop() interface{} {
	tailIndex := h.Len() - 1
	tail := (*h)[tailIndex]
	*h = (*h)[:tailIndex]
	return tail
}
