// re-implementation of C++ deleteDuplicates() in go
// Deletes duplicates from an array using a modified heap sort

package deldupes

import (
	"container/heap"
)

//  From the example Go heap package code
//  There must be a better to do this
type HeapHelper []int

func (self HeapHelper) Len() int           { return len(self) }
func (self HeapHelper) Less(i, j int) bool { return self[j] < self[i] } // max heap
func (self HeapHelper) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

// Acutally not used but interface demands it
func (self *HeapHelper) Push(x interface{}) { *self = append(*self, x.(int)) }

// Nota bene:  this is not the heap pop but a pop on the container
func (self *HeapHelper) Pop() interface{} {
	old := *self
	n := len(old) - 1
	x := old[n]
	*self = old[0:n]
	return x
}

func DeleteDuplicates(theArray []int) []int {
	// Need to get a mutable copy of the heap
	// There must be a better way
	heapHelper := make(HeapHelper, len(theArray))
	i := 0
	for _, value := range theArray {
		heapHelper[i] = value
		i++
	}
	heap.Init(&heapHelper)

	//  How to do this in-place?
	result := make([]int, 0)

	if len(theArray) > 0 {
		prevElement := heap.Pop(&heapHelper).(int)

		result = append(result, prevElement)

		for heapHelper.Len() > 0 {
			currentElement := heap.Pop(&heapHelper).(int)
			if currentElement != prevElement {
				result = append(result, currentElement)
				prevElement = currentElement
			}
		}
	}

	return result
}
