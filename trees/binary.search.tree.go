package trees

import "fmt"

type Comparable interface {
	Compare(c Comparable) int
}

type bstNode struct {
	left 	*bstNode
	right 	*bstNode
	parent	*bstNode
	value	Comparable
	level	uint16
}

func (bn *bstNode) calcLevel() {
	// Create local var
	l := uint16(0)
	if bn.left == nil && bn.right == nil {
		// Empty case just take zero
	} else if bn.left == nil {
		l = bn.right.level + 1
	} else if bn.right == nil {
		l = bn.left.level + 1
	} else {
		if bn.right.level >= bn.left.level {
			l = bn.right.level + 1
		} else {
			l = bn.left.level + 1
		}
	}
	// Replace in memory if changed
	if bn.level != l {
		bn.level = l
	}
}

type AVLTree struct {
	root *bstNode
}

func (a *AVLTree) Add(item Comparable)  {
	// If root was empty
	if a.root == nil {
		a.root = &bstNode{
			value:  item,
		}
		return
	}
	// Full add
	var node *bstNode
	root := a.root
	for {
		cmp := item.Compare(root.value)
		if cmp < 0 {
			if root.left == nil {
				root.left = &bstNode{
					value:  item,
				}
				node = root.left
				break
			}
			root = root.left
		} else if cmp > 0 {
			if root.right == nil {
				root.right = &bstNode{
					value:  item,
				}
				node = root.right
				break
			}
			root = root.right
		}
	}
	a.rebalance(node)
}

func (a *AVLTree) Delete(item interface{})  {

}

func (a *AVLTree) Search(item interface{})  {

}

func (a *AVLTree) Peek() (Comparable, error) {
	if a.root != nil {
		return a.root.value, nil
	}
	return nil, fmt.Errorf("empty")
}

func (a *AVLTree) PrettyPrint() string {
	return ""
}

func (a *AVLTree) rebalance(node *bstNode)  {
	for node != nil {
		node.calcLevel()
		lh := 0
		rh := 0
		if node.left != nil {
			lh = int(node.left.level)
		}
		if node.right != nil {
			rh = int(node.right.level)
		}
		sideCalc := lh - rh
		if sideCalc > 1 {
			a.rebalanceToRight(node)
		} else if sideCalc < -1 {
			a.rebalanceToLeft(node)
		}
		node = node.parent
	}
}

func (a *AVLTree) rebalanceToRight(node *bstNode) {
	var leftLeft int
	var leftRight int
	if node.left != nil && node.left.left != nil {
		leftLeft = int(node.left.left.level)
	}
	if node.right != nil && node.left.right != nil {
		leftRight = int(node.left.right.level)
	}
	if leftLeft >= leftRight {
		a.srr(node)
	} else {

	}
}

func (a *AVLTree) rebalanceToLeft(node *bstNode)  {
	var rightRight int
	var rightLeft int
	if node.right != nil && node.right.right != nil {
		rightRight = int(node.right.right.level)
	}
	if node.right != nil && node.right.left != nil {
		rightLeft = int(node.right.left.level)
	}
	if rightRight >= rightLeft {
		a.slr(node)
	} else {

	}
}

func (a *AVLTree) blr(node *bstNode)  {
	rotateChild := node.right
	a.slr(node)
	a.slr(node)
	a.srr(rotateChild)
}

func (a *AVLTree) brr(node *bstNode)  {
	rotateChild := node.left
	a.srr(node)
	a.srr(node)
	a.slr(rotateChild)
}

func (a *AVLTree) slr(prevRoot *bstNode)  {
	// Define rotation elements
	rotateCenter := prevRoot.right
	left := rotateCenter.left
	// Remap that previous root will be on the left of rotation center
	prevRoot.right = left
	rotateCenter.left = prevRoot
	rotateCenter.parent = prevRoot.parent
	prevRoot.parent = rotateCenter
	// Check if rotator left node did exist and switch it's parent
	if left != nil {
		left.parent = prevRoot
	}
	// Update the metadata on nodes
	prevRoot.calcLevel()
	rotateCenter.calcLevel()
	// Update tree root if prevRoot was actual tree root
	if prevRoot == a.root {
		a.root = rotateCenter
		return
	}
	// If prevRoot was not the tree root then update parent node
	// depending on which branch the prevRoot did reside
	if rotateCenter.parent.left == prevRoot {
		rotateCenter.parent.left = rotateCenter
	} else {
		rotateCenter.parent.right = rotateCenter
	}
}

func (a *AVLTree) srr(prevRoot *bstNode)  {
	// Define rotation elements
	rotateCenter := prevRoot.left
	right := rotateCenter.right
	// Remap that previous root will be on the left of rotation center
	prevRoot.left = right
	rotateCenter.right = prevRoot
	rotateCenter.parent = prevRoot.parent
	prevRoot.parent = rotateCenter
	// Check if rotator left node did exist and switch it's parent
	if right != nil {
		right.parent = prevRoot
	}
	// Update the metadata on nodes
	prevRoot.calcLevel()
	rotateCenter.calcLevel()
	// Update tree root if prevRoot was actual tree root
	if prevRoot == a.root {
		a.root = rotateCenter
		return
	}
	// If prevRoot was not the tree root then update parent node
	// depending on which branch the prevRoot did reside
	if rotateCenter.parent.left == prevRoot {
		rotateCenter.parent.left = rotateCenter
	} else {
		rotateCenter.parent.right = rotateCenter
	}
}
