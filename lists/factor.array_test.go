package lists

import (
	"fmt"
	"testing"
)

func TestFactorArray_AddStruct(t *testing.T) {

	a := NewFactorArray(nil)

	a.Add(&TestStruct{name: "one"})
	a.Add(&TestStruct{name: "two"})
	a.Add(&TestStruct{name: "three"})

	size := a.Size()
	fmt.Printf("\nstruct array size %d", size)

	one := &TestStruct{}
	a.GetValue(0, one)
	two := &TestStruct{}
	a.GetValue(1, two)
	three := &TestStruct{}
	a.GetValue(2, three)

	fmt.Printf("\nstruct values received [%+v] [%+v] [%+v]", one, two, three)

	remOne := &TestStruct{}
	a.Remove(0, remOne)
	remTwo := &TestStruct{}
	a.Remove(1, remTwo)
	remThree := &TestStruct{}
	a.Remove(0, remThree)

	fmt.Printf("\nstruct values deleted [%+v] [%+v] [%+v]", remOne, remTwo, remThree)

	size = a.Size()
	fmt.Printf("\nstruct array size %d", size)
}

func TestFactorArray_AddString(t *testing.T) {

	a := NewFactorArray(nil)

	a.Add("one")
	a.Add("two")
	a.Add("three")

	size := a.Size()
	fmt.Printf("\nstring array size %d", size)

	one := ""
	a.GetValue(0, &one)
	two := ""
	a.GetValue(1, &two)
	three := ""
	a.GetValue(2, &three)

	fmt.Printf("\nstring values received [%+v] [%+v] [%+v]", one, two, three)

	remOne := ""
	a.Remove(0, &remOne)
	remTwo := ""
	a.Remove(1, &remTwo)
	remThree := ""
	a.Remove(0, &remThree)

	fmt.Printf("\nvalues deleted [%+v] [%+v] [%+v]", remOne, remTwo, remThree)

	size = a.Size()
	fmt.Printf("\narray size %d", size)
}

func TestFactorArray_AddInt(t *testing.T) {

	a := NewFactorArray(nil)

	a.Add(1)
	a.Add(2)
	a.Add(3)

	size := a.Size()
	fmt.Printf("\nstring array size %d", size)

	one := 0
	a.GetValue(0, &one)
	two := 0
	a.GetValue(1, &two)
	three := 0
	a.GetValue(2, &three)

	fmt.Printf("\nstring values received [%+v] [%+v] [%+v]", one, two, three)

	remOne := 0
	a.Remove(0, &remOne)
	remTwo := 0
	a.Remove(1, &remTwo)
	remThree := 0
	a.Remove(0, &remThree)

	fmt.Printf("\nvalues deleted [%+v] [%+v] [%+v]", remOne, remTwo, remThree)

	size = a.Size()
	fmt.Printf("\narray size %d", size)
}

func TestFactorArray_Fill(t *testing.T) {
	a := NewFactorArray(nil)
	for i:=0; i<1000; i++ {
		a.Add(i)
	}
	fmt.Printf("\narray size after 1000 entries %d", a.size)
	for i:=0; i<999; i++ {
		a.Remove(1, nil)
	}
	a.Remove(0, nil)
	fmt.Printf("\narray size after deletion of 1000 entries %d", a.size)
}
