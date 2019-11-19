package lists

import (
	"fmt"
	"testing"
)


func TestLinkedListOrdered_AddStruct(t *testing.T) {

	sa := &LinkedListOrdered{comparator:OrderedIntComparator}

	sa.Add(&TestStruct{name: "one"}, 3)
	sa.Add(&TestStruct{name: "two"}, 2)
	sa.Add(&TestStruct{name: "three"}, 1)

	size := sa.Size()
	fmt.Printf("\nstruct array size %d", size)

	one := &TestStruct{}
	sa.GetValue(0, one)
	two := &TestStruct{}
	sa.GetValue(1, two)
	three := &TestStruct{}
	sa.GetValue(2, three)

	fmt.Printf("\nstruct values received [%+v] [%+v] [%+v]", one, two, three)

	remOne := &TestStruct{}
	sa.Remove(0, remOne)
	remTwo := &TestStruct{}
	sa.Remove(1, remTwo)
	remThree := &TestStruct{}
	sa.Remove(0, remThree)

	fmt.Printf("\nstruct values deleted [%+v] [%+v] [%+v]", remOne, remTwo, remThree)

	size = sa.Size()
	fmt.Printf("\nstruct array size %d", size)
}

func TestLinkedListOrdered_AddString(t *testing.T) {

	sa := &LinkedListOrdered{comparator: OrderedIntComparator}

	sa.Add("one", 1)
	sa.Add("two", 2)
	sa.Add("three", 3)

	size := sa.Size()
	fmt.Printf("\nstring array size %d", size)

	one := ""
	sa.GetValue(0, &one)
	two := ""
	sa.GetValue(1, &two)
	three := ""
	sa.GetValue(2, &three)

	fmt.Printf("\nstring values received [%+v] [%+v] [%+v]", one, two, three)

	remOne := ""
	sa.Remove(0, &remOne)
	remTwo := ""
	sa.Remove(1, &remTwo)
	remThree := ""
	sa.Remove(0, &remThree)

	fmt.Printf("\nvalues deleted [%+v] [%+v] [%+v]", remOne, remTwo, remThree)

	size = sa.Size()
	fmt.Printf("\narray size %d", size)
}

func TestLinkedListOrdered_AddInt(t *testing.T) {

	sa := &LinkedListOrdered{comparator:OrderedIntComparator}

	sa.Add(1, -1)
	sa.Add(2, 0)
	sa.Add(3, 1)

	size := sa.Size()
	fmt.Printf("\nstring array size %d", size)

	one := 0
	sa.GetValue(0, &one)
	two := 0
	sa.GetValue(1, &two)
	three := 0
	sa.GetValue(2, &three)

	fmt.Printf("\nstring values received [%+v] [%+v] [%+v]", one, two, three)

	remOne := 0
	sa.Remove(0, &remOne)
	remTwo := 0
	sa.Remove(1, &remTwo)
	remThree := 0
	sa.Remove(0, &remThree)

	fmt.Printf("\nvalues deleted [%+v] [%+v] [%+v]", remOne, remTwo, remThree)

	size = sa.Size()
	fmt.Printf("\narray size %d", size)
}

func TestLinkedListOrdered_Fill(t *testing.T) {
	sa := &LinkedListOrdered{comparator: OrderedIntComparator}
	for i:=0; i<1000; i++ {
		sa.Add(i, i % 10)
	}
	fmt.Printf("\narray size after 1000 entries %d", sa.size)
	for i:=0; i<999; i++ {
		sa.Remove(1, nil)
	}
	sa.Remove(0, nil)
	fmt.Printf("\narray size after deletion of 1000 entries %d", sa.size)
}
