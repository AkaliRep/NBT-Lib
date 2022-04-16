package types

import "fmt"

type NBTDouble struct {
	payload float64
}

func (nbtDouble NBTDouble) Payload() float64 { return nbtDouble.payload }
func (nbtDouble NBTDouble) String() string   { return fmt.Sprintf("%f", nbtDouble.payload) }

func NewNbtDouble(val float64) NBTDouble {
	return NBTDouble{
		payload: val,
	}
}

var _ NBTTag[float64] = NBTDouble{}
