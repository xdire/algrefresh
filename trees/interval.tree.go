package trees

import "math"

type IntervalSumTree struct {
	tree 		[]float64
}

func NewIntervalTreeFromSlice(from []float64) *IntervalSumTree {
	// Build the tree where
	newSize := 1 << uint64(math.Log(float64(len(from) - 1) + 1))
	it := &IntervalSumTree{tree: make([]float64, newSize)}
	it.buildTree(from)
	return it
}

func (it *IntervalSumTree) buildTree(from []float64) {
	elements := len(from)
	for i := elements; i < len(it.tree); i++ {
		it.tree[i] = from[i - elements]
	}
	for i := elements - 1; i > 0; i-- {
		it.tree[i] = it.tree[i * 2] + it.tree[(i * 2) + 1]
	}
}

func (it *IntervalSumTree) SumAt(from, to int) {

}