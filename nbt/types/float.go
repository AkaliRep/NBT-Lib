package types

import "fmt"

type NBTFloat struct {
	payload float32
}

func (nbtFloat NBTFloat) Payload() float32 { return nbtFloat.payload }
func (nbtFloat NBTFloat) String() string   { return fmt.Sprintf("%f", nbtFloat.payload) }

func NewNbtFloat(val float32) NBTFloat {
	return NBTFloat{
		payload: val,
	}
}

var _ NBTTag[float32] = NBTFloat{}
