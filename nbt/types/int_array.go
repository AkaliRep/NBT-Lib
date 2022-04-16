package types

import "fmt"

type NBTIntArray struct {
	payload []int32
}

func (nbtIntArray NBTIntArray) Payload() []int32 { return nbtIntArray.payload }
func (nbtIntArray NBTIntArray) String() string   { return fmt.Sprintf("%v", nbtIntArray.payload) }

func NewNbtIntArray(val []int32) NBTIntArray {
	return NBTIntArray{
		payload: val,
	}
}

var _ NBTTag[[]int32] = NBTIntArray{}
