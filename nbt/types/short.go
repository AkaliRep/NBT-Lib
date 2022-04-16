package types

import "fmt"

type NBTShort struct {
	payload int16
}

func (short NBTShort) Payload() int16 { return short.payload }
func (short NBTShort) String() string { return fmt.Sprintf("%d", short.payload) }

func NewNbtShort(val int16) NBTShort {
	return NBTShort{
		payload: val,
	}
}

var _ NBTTag[int16] = NBTShort{}
