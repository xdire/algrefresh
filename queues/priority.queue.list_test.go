package queues

import (
	"fmt"
	"github.com/xdire/algrefresh/lists"
	"testing"
)

type testStruct struct {
	name string
}

func TestPriorityQueueList_AddStruct(t *testing.T) {

	sa := &PriorityQueueList{}
	sa.list = lists.NewLinkedListOrdered(lists.OrderedIntComparator)

	sa.Add(&testStruct{name: "three"}, 3)
	sa.Add(&testStruct{name: "two"}, 2)
	sa.Add(&testStruct{name: "one"}, 1)
	sa.Add(&testStruct{name: "three three"}, 3)
	sa.Add(&testStruct{name: "two two"}, 2)
	sa.Add(&testStruct{name: "one one"}, 1)

	size := sa.Size()
	fmt.Printf("\nstruct array size %d", size)

	one := &testStruct{}
	err := sa.PollValue(one)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	two := &testStruct{}
	err = sa.PollValue(two)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	three := &testStruct{}
	err = sa.PollValue(three)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Printf("\nstruct values received [%+v] [%+v] [%+v]", one, two, three)

	one = &testStruct{}
	err = sa.PollValue(one)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	two = &testStruct{}
	err = sa.PollValue(two)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	three = &testStruct{}
	err = sa.PollValue(three)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Printf("\nstruct values deleted [%+v] [%+v] [%+v]", one, two, three)

	size = sa.Size()
	fmt.Printf("\nstruct array size %d", size)
}

func TestPriorityQueueList_AddString(t *testing.T) {

	sa := &PriorityQueueList{}
	sa.list = lists.NewLinkedListOrdered(lists.OrderedIntComparator)

	sa.Add("one", 1)
	sa.Add("two", 2)
	sa.Add("three", 3)
	sa.Add("three three", 3)
	sa.Add("two two", 2)
	sa.Add("one one", 1)

	size := sa.Size()
	fmt.Printf("\nstring pq size %d", size)

	one := ""
	err := sa.PollValue(&one)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	two := ""
	err = sa.PollValue(&two)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	three := ""
	err = sa.PollValue(&three)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Printf("\nstring values received [%+v] [%+v] [%+v]", one, two, three)

	one = ""
	err = sa.PollValue(&one)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	two = ""
	err = sa.PollValue(&two)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	three = ""
	err = sa.PollValue(&three)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Printf("\nvalues deleted [%+v] [%+v] [%+v]", one, two, three)

	size = sa.Size()
	fmt.Printf("\nstring pq size %d", size)
}

func TestPriorityQueueList_Fill(t *testing.T) {
	sa := &PriorityQueueList{}
	sa.list = lists.NewLinkedListOrdered(lists.OrderedIntComparator)
	for i:=0; i<1000; i++ {
		sa.Add(i, i % 10)
	}
	fmt.Printf("\npq size after 1000 entries %d", sa.size)
	for i:=0; i<999; i++ {
		err := sa.PollValue(nil)
		if err != nil {
			t.Fail()
		}
	}
	err := sa.PollValue(nil)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("\npq size after polling of 1000 entries %d", sa.size)
}
