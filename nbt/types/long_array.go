package types

import "fmt"

type NBTLongArray struct {
	payload []int64
}

func (nbtLongArray NBTLongArray) Payload() []int64 { return nbtLongArray.payload }
func (nbtLongArray NBTLongArray) String() string   { return fmt.Sprintf("%v", nbtLongArray.payload) }

func NewNbtLongArray(val []int64) NBTLongArray {
	return NBTLongArray{
		payload: val,
	}
}

var _ NBTTag[[]int64] = NBTLongArray{}
