package fibheap

// Example Test for the Fibonacci heap with floats

import (
	"fmt"
)

func Example() {
	heap1 := NewFloatFibHeap()
	fmt.Println("Created heap 1.")
	fmt.Printf("Heap 1 insert: %v\n", heap1.Enqueue(15.77))
	heap2 := NewFloatFibHeap()
	fmt.Println("Created heap 2.")
	fmt.Printf("Heap 2 is empty? %v\n", heap2.IsEmpty())
	fmt.Printf("Heap 2 insert: %v\n", heap2.Enqueue(-1002.2001))
	fmt.Printf("Heap 2 insert: %v\n", heap2.Enqueue(-0.001))
	fmt.Printf("Heap 1 size: %v\n", heap1.Size())
	fmt.Printf("Heap 2 size: %v\n", heap2.Size())
	fmt.Printf("Heap 1 is empty? %v\n", heap1.IsEmpty())
	fmt.Printf("Heap 2 is empty? %v\n", heap2.IsEmpty())

	fmt.Printf("\nMerge Heap 1 and Heap 2.\n")
	mergedHeap := heap1.Merge(&heap2)
	fmt.Printf("Merged heap size: %v\n", mergedHeap.Size())
	fmt.Printf("Set node with priority %v to new priority %v\n", -1002.2001, mergedHeap.DecreaseKey(-1002.2001, -1003.0))
	fmt.Printf("Dequeue minimum of merged heap: %v\n", mergedHeap.DequeueMin())
	fmt.Printf("Merged heap size: %v\n", mergedHeap.Size())
	fmt.Printf("Delete from merged heap: %v\n", mergedHeap.Delete(-0.001))
	fmt.Printf("Merged heap size: %v\n", mergedHeap.Size())
	fmt.Printf("Extracting minimum of merged heap: %v\n", mergedHeap.DequeueMin())
	fmt.Printf("Merged heap size: %v\n", mergedHeap.Size())
	fmt.Printf("Merged heap is empty? %v\n", mergedHeap.IsEmpty())

	// Output:
	// Created heap 1.
	// Heap 1 insert: 15.77
	// Created heap 2.
	// Heap 2 is empty? Yes
	// Heap 2 insert: -1002.2001
	// Heap 2 insert: -0.001
	// Heap 1 size: 1
	// Heap 2 size: 2
	// Heap 1 is empty? No
	// Heap 2 is empty? No
	//
	// Merge Heap 1 and Heap 2.
	// Merged heap size: 3
	// Set node with priority -1002.2001 to new priority -1003.0
	// Dequeue minimum of merged heap: -1003.0
	// Merged heap size: 2
	// Delete from merged heap: -0.001
	// Merged heap size: 1
	// Extracting minimum of merged heap: 15.77
	// Merged heap size: 0
	// Merged heap is empty? Yes
}
