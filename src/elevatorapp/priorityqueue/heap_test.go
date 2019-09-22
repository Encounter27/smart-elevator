package priorityqueue

import (
	"container/heap"
	"testing"
)

func TestMinHeapWithItem(t *testing.T) {
	items := map[string]int{
		"banana": 5, "apple": 4, "pear": 3,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.value, item.priority)

	out := []string{"orange", "pear", "apple", "banana"}
	k := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.value != out[k] {
			t.Errorf("Incorrect min heap order: expected:- %s, got :- %s", out[k], item.value)
		}
		k++
	}
}

func TestMinHeap(t *testing.T) {
	items := []int{4, 3, 2}

	pq := make(PriorityQueue, len(items))
	i := 0
	for _, priority := range items {
		pq[i] = &Item{
			value:    nil,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := &Item{
		value:    nil,
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.value, item.priority)

	out := []int{1, 2, 3, 4}
	k := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.priority != out[k] {
			t.Errorf("Incorrect min heap order: expected:- %v, got :- %v", out[k], item.priority)
		}
		k++
	}
}
