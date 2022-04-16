package types

type NBTString struct {
	payload string
}

func (nbtString NBTString) Payload() string { return nbtString.payload }
func (nbtString NBTString) String() string  { return nbtString.payload }

func NewNbtString(val string) NBTString {
	return NBTString{
		payload: val,
	}
}

var _ NBTTag[string] = NBTString{}
