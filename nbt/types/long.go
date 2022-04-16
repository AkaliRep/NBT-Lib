package types

import "fmt"

type NBTLong struct {
	payload int64
}

func (nbtLong NBTLong) Payload() int64 { return nbtLong.payload }
func (nbtLong NBTLong) String() string { return fmt.Sprintf("%d", nbtLong.payload) }

func NewNbtLong(val int64) NBTLong {
	return NBTLong{
		payload: val,
	}
}

var _ NBTTag[int64] = NBTLong{}
