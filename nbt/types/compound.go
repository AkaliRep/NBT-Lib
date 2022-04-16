package types

import (
	"errors"
	"fmt"
)

type NBTCompound struct {
	name    string
	payload map[string]NBTTag[any]
}

func (compound NBTCompound) Name() string                    { return compound.name }
func (compound NBTCompound) Payload() map[string]NBTTag[any] { return compound.payload }
func (compound NBTCompound) String() string {
	return fmt.Sprintf("NBTCompound(%s, %s)", compound.name, compound.payload)
}

func (compound NBTCompound) GetByte(s string) (NBTByte, error) {
	tag := compound.payload[s]
	if t, ok := tag.(NBTByte); ok {
		return t, nil
	}

	return NBTByte{}, errors.New("")
}

func NewNbtNamedCompound(name string) NBTCompound {
	return NBTCompound{
		name:    name,
		payload: make(map[string]NBTTag[any]),
	}
}

func NewNbtCompound() NBTCompound {
	return NBTCompound{
		payload: make(map[string]NBTTag[any]),
	}
}

var _ NBTTag[map[string]NBTTag[any]] = NBTCompound{}
