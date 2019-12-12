package hw01_09_avl_tree

import (
	"fmt"
	"github.com/xdire/algrefresh/trees"
	"testing"
)

type data int16

func (d data) Compare(cmpTo trees.Comparable) int {
	if dd, ok := cmpTo.(data); ok {
		if d > dd {
			return 1
		} else if d < dd {
			return -1
		}
	}
	return 0
}

func TestAVLTree_Add(t *testing.T) {
	tree := &trees.AVLTree{}
	type args struct {
		items []trees.Comparable
	}
	tests := []struct {
		name   string
		args   args
	}{
		{
			"Testing the AVL tree addition",
			args{items: []trees.Comparable{
				data(4), data(1), data(3), data(6), data(-1), data(15), data(123),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.items {
				tree.Add(v)
			}
			root, err := tree.Peek()
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("%+v", root)
		})
	}
}