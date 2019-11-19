package lists

import (
	"fmt"
	"testing"
)

func TestSingleArray_AddStruct(t *testing.T) {

	sa := &SingleArray{}

	sa.Add(&TestStruct{name: "one"})
	sa.Add(&TestStruct{name: "two"})
	sa.Add(&TestStruct{name: "three"})

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

func TestSingleArray_AddString(t *testing.T) {

	sa := &SingleArray{}

	sa.Add("one")
	sa.Add("two")
	sa.Add("three")

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

func TestSingleArray_AddInt(t *testing.T) {

	sa := &SingleArray{}

	sa.Add(1)
	sa.Add(2)
	sa.Add(3)

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
