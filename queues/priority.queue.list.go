package queues

import (
	"errors"
	"github.com/xdire/algrefresh/lists"
)

type PriorityQueueList struct {
	list 	*lists.LinkedListOrdered
	size 	uint32
}

func NewPriorityQueueList(comparator lists.AbstractComparator) *PriorityQueueList {
	return &PriorityQueueList{
		list: lists.NewLinkedListOrdered(comparator),
		size: 0,
	}
}

func (p *PriorityQueueList) SetComparator(c func(a interface{}, b interface{}) (int8, error)) {
	p.list.SetComparator(c)
}

func (p *PriorityQueueList) Add(item interface{}, order interface{}) {
	ll := &lists.LinkedList{}
	llr := p.list.Once(ll, order).(*lists.LinkedList)
	llr.Add(item)
	p.size++
}

func (p *PriorityQueueList) Poll() (interface{}, error) {
	return p.pPoll(nil)
}

func (p *PriorityQueueList) Peek() (interface{}, error) {
	return p.pPeek(nil)
}

func (p *PriorityQueueList) PollValue(receiver interface{}) error {
	_, err := p.pPoll(receiver)
	return err
}

func (p *PriorityQueueList) PeekValue(receiver interface{}) error {
	_, err := p.pPeek(receiver)
	return err
}

func (p *PriorityQueueList) Size() uint32 {
	return p.size
}

func (p *PriorityQueueList) pPoll(receiver interface{}) (interface{}, error) {
	llr, err := p.list.Peek()
	if err != nil {
		return nil, err
	}
	ll := llr.(*lists.LinkedList)
	// RemoveValue first index of sublist if size at least 2
	if ll.Size() > 1 {
		p.size--
		if receiver != nil {
			ll.RemoveValue(0, receiver)
			return nil, nil
		}
		return ll.Remove(0)
	} else if ll.Size() == 1 {
	// RemoveValue first index and remove the order
	// altogether if size is 1
		if receiver != nil {
			ll.RemoveValue(0, receiver)
			_, err = p.list.Poll()
			p.size--
			return nil, nil
		}
		v, err := ll.Remove(0)
		if err != nil {
			return nil, err
		}
		_, err = p.list.Poll()
		p.size--
		return v, err
	} else {
	// If by somehow attached linked list was zeroed
	// in size, so we can't get any information from that
	// remove it and repeat
		_, err = p.list.Poll()
		return p.pPoll(receiver)
	}
}

func (p *PriorityQueueList) pPeek(receiver interface{}) (interface{}, error) {
	llr, err := p.list.Peek()
	if err != nil {
		return nil, err
	}
	ll := llr.(*lists.LinkedList)
	// If size is positive get first
	if ll.Size() > 0 {
		if receiver != nil {
			ll.GetValue(0, receiver)
			return nil, nil
		}
		return ll.Get(0)
	}
	// Should never happen due proper removal but covered
	return nil, errors.New("queue malfunction")
}
