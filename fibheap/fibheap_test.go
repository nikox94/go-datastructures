package fibheap

// Example tests for the Fibonacci hea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImmutableSingleDimensionAdd(t *testing.T) {
	tree := newImmutableRangeTree(1)
	entry := constructMockEntry(0, int64(0), int64(0))
	tree2 := tree.Add(entry)

	result := tree.Query(
		constructMockInterval(dimension{0, 10}, dimension{0, 10}),
	)
	assert.Len(t, result, 0)

	result = tree2.Query(
		constructMockInterval(dimension{0, 10}, dimension{0, 10}),
	)
	assert.Equal(t, Entries{entry}, result)

	// // // // // // // // // //

	heap := newFibHeap( /*expected number of elements here*/ )
	heap.Add(newFibHeap() /*elem*/, 0 /*priority*/)
	heap.Min()
	heap.IsEmpty()
	heap.Size()

	heap1 := newFibHeap() /*Some heap*/
	heap2 := newFibHeap() /*Other heap*/
	heap = Merge(heap1, heap2)

	heap.DelMin()
	heap.DecreaseKey(newFibHeap() /*elem*/, 0 /*priority*/)
	heap.Delete(newFibHeap() /*entry*/)
}

func BenchmarkInsertWithExpand(b *testing.B) {
	numItems := uint64(1000)

	hms := make([]*FastIntegerHashMap, 0, b.N)
	for i := 0; i < b.N; i++ {
		hm := New(10)
		hms = append(hms, hm)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hm := hms[i]
		for j := uint64(0); j < numItems; j++ {
			hm.Set(j, j)
		}
	}
}
