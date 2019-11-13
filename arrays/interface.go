package arrays

type IArray interface {
	Size() uint32
	Add(item interface{})
	Get(index uint32, receiver interface{})
	Remove(index uint32, receiver interface{})
	AddAtIndex(index uint32, item interface{})
}
