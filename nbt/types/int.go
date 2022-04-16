package types

import "fmt"

type NBTInt struct {
	payload int32
}

func (nbtInt NBTInt) Payload() int32 { return nbtInt.payload }
func (nbtInt NBTInt) String() string { return fmt.Sprintf("%d", nbtInt.payload) }

func NewNbtInt(val int32) NBTInt {
	return NBTInt{
		payload: val,
	}
}

var _ NBTTag[int32] = NBTInt{}
