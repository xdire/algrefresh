package algrefresh

type IList interface {
	Size() uint32
	Add(item interface{})
	Get(index uint32) (interface{}, error)
	Remove(index uint32) (interface{}, error)
	GetValue(index uint32, receiver interface{})
	RemoveValue(index uint32, receiver interface{})
	AddAtIndex(index uint32, item interface{})
}

type IOrderedList interface {
	Size() uint32
	SetComparator(c func(a interface{}, b interface{}) (int8, error))
	Add(item interface{}, order interface{})
	Once(item interface{}, order interface{}) interface{}
	Get(index uint32) (interface{}, error)
	GetOrder(order interface{}) (interface{}, error)
	GetValue(index uint32, receiver interface{})
	GetOrderValue(order interface{}, receiver interface{}) error
	Remove(index uint32) (interface{}, error)
	RemoveValue(index uint32, receiver interface{})
}

type IQueue interface {
	Push(item interface{})
	Poll() (interface{}, error)
	Peek() (interface{}, error)
	PollValue(receiver interface{}) error
	PeekValue(receiver interface{}) error
	Size() uint32
}

type IPriorityQueue interface {
	SetComparator(c func(a interface{}, b interface{}) (int8, error))
	Add(item interface{}, order interface{})
	Poll() (interface{}, error)
	Peek() (interface{}, error)
	PollValue(receiver interface{}) error
	PeekValue(receiver interface{}) error
	Size() uint32
}
