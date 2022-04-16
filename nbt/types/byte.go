package types

import "fmt"

type NBTByte struct {
	payload byte
}

func (nbtByte NBTByte) Payload() byte  { return nbtByte.payload }
func (nbtByte NBTByte) String() string { return fmt.Sprintf("%d", nbtByte.payload) }

func NewNbtByte(val byte) NBTByte {
	return NBTByte{
		payload: val,
	}
}

var _ NBTTag[byte] = NBTByte{}
