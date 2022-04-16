package types

import "fmt"

type NBTByteArray struct {
	payload []byte
}

func (nbtByteArray NBTByteArray) Payload() []byte { return nbtByteArray.payload }
func (nbtByteArray NBTByteArray) String() string  { return fmt.Sprintf("%v", nbtByteArray.payload) }

func NewNbtByteArray(val []byte) NBTByteArray {
	return NBTByteArray{
		payload: val,
	}
}

var _ NBTTag[[]byte] = NBTByteArray{}
