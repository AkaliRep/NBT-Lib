package types

type NBTTag[T any] interface {
	Payload() T
	String() string
}
