package structures

type BinaryIndexedTree struct {
	tree     []int
	capacity int
}

func (bit *BinaryIndexedTree) Init(capacity int) {
	bit.tree, bit.capacity = make([]int, capacity+1), capacity
}

func (bit *BinaryIndexedTree) Add(index int, val int) {
	for ; index <= bit.capacity; index += index & -index {
		bit.tree[index] += val
	}
}

func (bit *BinaryIndexedTree) Query(index int) int {
	sum := 0
	for ; index > 0; index -= index & -index {
		sum += bit.tree[index]
	}
	return sum
}
