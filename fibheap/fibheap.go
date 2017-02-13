/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Special thanks to  Keith Schwarz (htiek@cs.stanford.edu),
whose code and documentation have been used as a reference
for the algorithm implementation.
http://www.keithschwarz.com/interesting/code/?dir=fibonacci-heap
*/

/*
* An implementation of a priority queue backed by a Fibonacci heap,
* as described by Fredman and Tarjan.  Fibonacci heaps are interesting
* theoretically because they have asymptotically good runtime guarantees
* for many operations.  In particular, insert, peek, and decrease-key all
* run in amortized O(1) time.  dequeueMin and delete each run in amortized
* O(lg n) time.  This allows algorithms that rely heavily on decrease-key
* to gain significant performance boosts.  For example, Dijkstra's algorithm
* for single-source shortest paths can be shown to run in O(m + n lg n) using
* a Fibonacci heap, compared to O(m lg n) using a standard binary or binomial
* heap.
*
* Internally, a Fibonacci heap is represented as a circular, doubly-linked
* list of trees obeying the min-heap property.  Each node stores pointers
* to its parent (if any) and some arbitrary child.  Additionally, every
* node stores its degree (the number of children it has) and whether it
* is a "marked" node.  Finally, each Fibonacci heap stores a pointer to
* the tree with the minimum value.
*
* To insert a node into a Fibonacci heap, a singleton tree is created and
* merged into the rest of the trees.  The merge operation works by simply
* splicing together the doubly-linked lists of the two trees, then updating
* the min pointer to be the smaller of the minima of the two heaps.  Peeking
* at the smallest element can therefore be accomplished by just looking at
* the min element.  All of these operations complete in O(1) time.
*
* The tricky operations are dequeueMin and decreaseKey.  dequeueMin works
* by removing the root of the tree containing the smallest element, then
* merging its children with the topmost roots.  Then, the roots are scanned
* and merged so that there is only one tree of each degree in the root list.
* This works by maintaining a dynamic array of trees, each initially null,
* pointing to the roots of trees of each dimension.  The list is then scanned
* and this array is populated.  Whenever a conflict is discovered, the
* appropriate trees are merged together until no more conflicts exist.  The
* resulting trees are then put into the root list.  A clever analysis using
* the potential method can be used to show that the amortized cost of this
* operation is O(lg n), see "Introduction to Algorithms, Second Edition" by
* Cormen, Rivest, Leiserson, and Stein for more details.
*
* The other hard operation is decreaseKey, which works as follows.  First, we
* update the key of the node to be the new value.  If this leaves the node
* smaller than its parent, we're done.  Otherwise, we cut the node from its
* parent, add it as a root, and then mark its parent.  If the parent was
* already marked, we cut that node as well, recursively mark its parent,
* and continue this process.  This can be shown to run in O(1) amortized time
* using yet another clever potential function.  Finally, given this function,
* we can implement delete by decreasing a key to -\infty, then calling
* dequeueMin to extract it.
 */

package fibheap

/******************************************
 ************** INTERFACE *****************
 ******************************************/

// The FloatingFibonacciHeap is an implementation of a fibonacci heap
// with only floating-point priorities and no user data attached.
type FloatingFibonacciHeap interface {
	// Adds and element to the heap
	Enqueue(priority float64) float64
	// Returns the minimum element in the heap
	Min() float64
	// Is the heap empty?
	IsEmpty() bool
	// The number of elements in the heap
	Size() uint
	// Removes and returns the minimal element
	// in the heap
	DequeueMin() float64
	// Decreases the key of the given element
	// and sets it to the new given priority
	// returns the new priority if succesfully set
	DecreaseKey(oldPriority, newPriority float64) float64
	// Deletes the given element in the heap
	// Will use the numeric value of the priority
	// as an identifier.
	// TODO: How do we deal with doubles?
	Delete(priority float64) float64
	// Will merge two heaps
	Merge(otherHeap *FloatingFibonacciHeap) FloatingFibonacciHeap
}

/******************************************
 ************** END INTERFACE *************
 ******************************************/

type fibHeap struct {
	min  *entry // The minimal element
	size uint
}

type entry struct {
	degree                    int
	isMarked                  bool
	next, prev, child, parent *entry
	priority                  float64
}

// ACTUAL CODE

func NewFloatFibHeap() FloatingFibonacciHeap { return &fibHeap{nil, 0} }

func (heap *fibHeap) Enqueue(priority float64) float64 {
	//result := newEntry(priority)

	// Merge singleton list with heap

	heap.size++
	return priority
}

func (heap *fibHeap) Min() float64 {
	if heap.IsEmpty() {
		return -42
	}
	return heap.min.priority
}

func (heap *fibHeap) IsEmpty() bool {
	return heap.size == 0
}

func (heap *fibHeap) Size() uint {
	return heap.size
}

func (heap *fibHeap) DequeueMin() float64 {
	return 42.0
}

func (heap *fibHeap) DecreaseKey(oldPriority, newPriority float64) float64 {
	return newPriority
}

func (heap *fibHeap) Delete(priority float64) float64 {
	return priority
}

/*
 * Given two Fibonacci heaps, returns a new Fibonacci heap that contains
 * all of the elements of the two heaps.  Each of the input heaps is
 * destructively modified by having all its elements removed.  You can
 * continue to use those heaps, but be aware that they will be empty
 * after this call completes.
 */
func (heap *fibHeap) Merge(other *FloatingFibonacciHeap) FloatingFibonacciHeap {
	//result := new(FibHeap)
	return NewFloatFibHeap()

	//result.size = fibHeap.size + other.size
	//result.min = fibHeap.mergeLists(other)
	//return result
}

/*func (fibHeap *FibHeap) delMin() *entry {
	if fibHeap.IsEmpty() {
		return nil
	}

	// TODO: fix
	defer func() { fibHeap.size-- }()
	minElem := fibHeap.min
	return nil
}

// HELPER FUNCTIONS

func newEntry(priority float64) entry {
	result := new(entry)
	result.next = &result
	result.prev = &result
	result.priority = priority
	return result
}*/
