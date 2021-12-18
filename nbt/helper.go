package nbt

import (
	"fmt"
	"reflect"
)

const (
	TAG_END byte = iota
	TAG_BYTE
	TAG_SHORT
	TAG_INT
	TAG_LONG
	TAG_FLOAT
	TAG_DOUBLE
	TAG_BYTE_ARRAY
	TAG_STRING
	TAG_LIST
	TAG_COMPOUND
	TAG_INT_ARRAY
	TAG_LONG_ARRAY
)

func TagTypeToString(tagType byte) string {
	switch tagType {
	case TAG_END:
		return "TAG_END"
	case TAG_BYTE:
		return "TAG_BYTE"
	case TAG_SHORT:
		return "TAG_SHORT"
	case TAG_INT:
		return "TAG_INT"
	case TAG_LONG:
		return "TAG_LONG"
	case TAG_FLOAT:
		return "TAG_FLOAT"
	case TAG_DOUBLE:
		return "TAG_DOUBLE"
	case TAG_BYTE_ARRAY:
		return "TAG_BYTE_ARRAY"
	case TAG_STRING:
		return "TAG_STRING"
	case TAG_LIST:
		return "TAG_LIST"
	case TAG_COMPOUND:
		return "TAG_COMPOUND"
	case TAG_INT_ARRAY:
		return "TAG_INT_ARRAY"
	case TAG_LONG_ARRAY:
		return "TAG_LONG_ARRAY"
	default:
		return "UNRECOGNIZED_TYPE"
	}
}

type NBTPayloadUnion struct {
	S8        byte
	S16       int16
	S32       int32
	S64       int64
	F32       float32
	F64       float64
	ByteArray []byte
	Str       string
	List      []NBTPayloadUnion
	Compound  map[string]NBTTag
	IntArray  []int32
	LongArray []int64
}

type NBTTag struct {
	NBTType    byte
	Name       string
	NBTPayload NBTPayloadUnion
}

type Chunk struct {
	X    int
	Y    int
	Data NBTTag
}

func PrintlAllSubTypes(t reflect.Type, deep int) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Type.Kind().String() == "struct" {
			PrintlAllSubTypes(f.Type, deep+1)
		}
		fmt.Printf("(%d, %s) -> %s\n", deep, f.Name, f.Type)
	}
}
