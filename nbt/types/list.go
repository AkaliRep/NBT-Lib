package types

import "fmt"

type NBTList[T any] struct {
	payload []T
}

func (nbtList NBTList[T]) Payload() []T   { return nbtList.payload }
func (nbtList NBTList[T]) String() string { return fmt.Sprintf("%v", nbtList.payload) }

func NewNbtList[T any](val []T) NBTList[T] {
	return NBTList[T]{
		payload: val,
	}
}

var _ NBTList[any] = NBTList[any]{}
