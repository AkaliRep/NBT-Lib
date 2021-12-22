package nbt

import (
	"fmt"
	"reflect"
	"strings"
)

var deep = 0

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

func (tag NBTTag) String() string {
	tabs := strings.Repeat("\t", deep)
	strFormat := tabs + "NBTTag(Type: %s, Name: %s, NBTPayload: %s)\n"
	intFormat := tabs + "NBTTag(Type: %s, Name: %s, NBTPayload: %d)\n"
	floatFormat := tabs + "NBTTag(Type: %s, Name: %s, NBTPayload: %f)\n"
	listFormat := tabs + "NBTTag(Type: %s, Name: %s, NBTPayload: %v)\n"
	compoundFormat := tabs + "NBTTag(Type: %s, Name: %s, NBTPayload)\n%s"
	payload := tag.NBTPayload
	str := ""

	switch tag.NBTType {
	case TAG_BYTE:
		str += fmt.Sprintf(intFormat, TagTypeToString(TAG_BYTE), tag.Name, payload.S8)
	case TAG_SHORT:
		str += fmt.Sprintf(intFormat, TagTypeToString(TAG_SHORT), tag.Name, payload.S16)
	case TAG_INT:
		str += fmt.Sprintf(intFormat, TagTypeToString(TAG_INT), tag.Name, payload.S32)
	case TAG_LONG:
		str += fmt.Sprintf(intFormat, TagTypeToString(TAG_LONG), tag.Name, payload.S64)
	case TAG_FLOAT:
		str += fmt.Sprintf(floatFormat, TagTypeToString(TAG_FLOAT), tag.Name, payload.F32)
	case TAG_DOUBLE:
		str += fmt.Sprintf(floatFormat, TagTypeToString(TAG_DOUBLE), tag.Name, payload.F64)
	case TAG_STRING:
		str += fmt.Sprintf(strFormat, TagTypeToString(TAG_STRING), tag.Name, payload.Str)
	case TAG_LIST:
		return fmt.Sprintf(listFormat, TagTypeToString(TAG_LIST), tag.Name, payload.List)
	case TAG_COMPOUND:
		result := ""
		deep += 1
		for k, v := range tag.NBTPayload.Compound {
			v.Name = k
			result += v.String()
		}

		deep -= 1
		str += fmt.Sprintf(compoundFormat, TagTypeToString(TAG_COMPOUND), tag.Name, result)
	case TAG_INT_ARRAY:
		str += fmt.Sprintf(listFormat, TagTypeToString(TAG_INT_ARRAY), tag.Name, payload.IntArray)
	default:
		panic("UNREACHEABLE")
	}

	return str
}

type Chunk struct {
	X    int
	Y    int
	Data NBTTag
}

func FillStruct(iface interface{}, deep int) {
	v := reflect.ValueOf(iface)
	t := reflect.TypeOf(iface)

	for i := 0; i < t.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)

		fmt.Printf("(%s, %d) -> %s\n", ft.Name, deep, ft.Type.Name())
		if a, ok := fv.Interface().(NBTTag); ok {
			fmt.Printf("%v\n", a)
		}

		if fv.Kind() == reflect.Struct {
			FillStruct(fv.Interface(), deep+1)
		}
	}
}
