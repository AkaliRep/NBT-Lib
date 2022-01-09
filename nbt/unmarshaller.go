package nbt

import (
	"encoding/binary"
	"math"
)

func WriteByte(output *[]byte, data byte) {
	*output = append(*output, data)
}

func WriteShortBE(output *[]byte, data int16) {
	*output = append(*output, byte(data>>8))
	*output = append(*output, byte(data))
}

func WriteIntBE(output *[]byte, data int32) {
	*output = append(*output, byte(data>>24))
	*output = append(*output, byte(data>>16))
	*output = append(*output, byte(data>>8))
	*output = append(*output, byte(data))
}

func WriteLongBE(output *[]byte, data int64) {
	*output = append(*output, byte(data>>56))
	*output = append(*output, byte(data>>48))
	*output = append(*output, byte(data>>40))
	*output = append(*output, byte(data>>32))
	*output = append(*output, byte(data>>24))
	*output = append(*output, byte(data>>16))
	*output = append(*output, byte(data>>8))
	*output = append(*output, byte(data))
}

func WriteFloatBE(output *[]byte, data float32) {
	t := math.Float32bits(data)
	r := make([]byte, 4)
	binary.BigEndian.PutUint32(r, t)
	*output = append(*output, r...)
}

func WriteDoubleBE(output *[]byte, data float64) {
	t := math.Float64bits(data)
	r := make([]byte, 8)
	binary.BigEndian.PutUint64(r, t)
	*output = append(*output, r...)
}

func WriteByteArray(output *[]byte, data []byte) {
	WriteIntBE(output, int32(len(data)))
	*output = append(*output, data...)
}

func WriteString(output *[]byte, data string) {
	WriteShortBE(output, int16(len(data)))
	*output = append(*output, []byte(data)...)
}

func WriteList(output *[]byte, list NBTList) {
	WriteByte(output, list.Type)
	switch list.Type {
	case TAG_BYTE:
		WriteByteArray(output, list.Data.S8)
	case TAG_SHORT:
		WriteIntBE(output, int32(len(list.Data.S16)))
		for i := 0; i < len(list.Data.S16); i++ {
			WriteShortBE(output, list.Data.S16[i])
		}

	case TAG_INT:
		WriteIntBE(output, int32(len(list.Data.S32)))
		for i := 0; i < len(list.Data.S32); i++ {
			WriteIntBE(output, list.Data.S32[i])
		}

	case TAG_LONG:
		WriteIntBE(output, int32(len(list.Data.S64)))
		for i := 0; i < len(list.Data.S64); i++ {
			WriteLongBE(output, list.Data.S64[i])
		}

	case TAG_FLOAT:
		WriteIntBE(output, int32(len(list.Data.F32)))
		for i := 0; i < len(list.Data.F32); i++ {
			WriteFloatBE(output, list.Data.F32[i])
		}

	case TAG_DOUBLE:
		WriteIntBE(output, int32(len(list.Data.F64)))
		for i := 0; i < len(list.Data.F64); i++ {
			WriteDoubleBE(output, list.Data.F64[i])
		}

	case TAG_BYTE_ARRAY:
		WriteIntBE(output, int32(len(list.Data.ByteArray)))
		for i := 0; i < len(list.Data.ByteArray); i++ {
			WriteByteArray(output, list.Data.ByteArray[i])
		}

	case TAG_STRING:
		WriteIntBE(output, int32(len(list.Data.Str)))
		for i := 0; i < len(list.Data.Str); i++ {
			WriteString(output, list.Data.Str[i])
		}

	case TAG_LIST:
		WriteIntBE(output, int32(len(list.Data.List)))
		for i := 0; i < len(list.Data.List); i++ {
			WriteList(output, list.Data.List[i])
		}

	case TAG_COMPOUND:
		WriteIntBE(output, int32(len(list.Data.Compound)))
		for i := 0; i < len(list.Data.Compound); i++ {
			WriteCompound(output, list.Data.Compound[i])
		}

	case TAG_INT_ARRAY:
		WriteIntBE(output, int32(len(list.Data.IntArray)))
		for i := 0; i < len(list.Data.IntArray); i++ {
			WriteIntArray(output, list.Data.IntArray[i])
		}

	case TAG_LONG_ARRAY:
		WriteIntBE(output, int32(len(list.Data.LongArray)))
		for i := 0; i < len(list.Data.LongArray); i++ {
			WriteLongArray(output, list.Data.LongArray[i])
		}
	}
}

func WriteIntArray(output *[]byte, data []int32) {
	WriteIntBE(output, int32(len(data)))
	for i := 0; i < len(data); i++ {
		WriteIntBE(output, data[i])
	}
}

func WriteLongArray(output *[]byte, data []int64) {
	WriteIntBE(output, int32(len(data)))
	for i := 0; i < len(data); i++ {
		WriteLongBE(output, data[i])
	}
}

func WriteCompound(output *[]byte, data map[string]NBTTag) {
	for k, v := range data {
		WriteByte(output, v.NBTType)
		WriteString(output, k)
		WriteTag(output, v)
	}
	WriteByte(output, TAG_END)
}

func WriteImplicitCompound(output *[]byte, data NBTTag) {
	WriteByte(output, data.NBTType)
	WriteString(output, data.Name)
	WriteCompound(output, data.NBTPayload.Compound)
}

func WriteTag(output *[]byte, tag NBTTag) {
	switch tag.NBTType {
	case TAG_BYTE:
		WriteByte(output, tag.NBTPayload.S8)

	case TAG_SHORT:
		WriteShortBE(output, tag.NBTPayload.S16)

	case TAG_INT:
		WriteIntBE(output, tag.NBTPayload.S32)

	case TAG_LONG:
		WriteLongBE(output, tag.NBTPayload.S64)

	case TAG_FLOAT:
		WriteFloatBE(output, tag.NBTPayload.F32)

	case TAG_DOUBLE:
		WriteDoubleBE(output, tag.NBTPayload.F64)

	case TAG_BYTE_ARRAY:
		WriteByteArray(output, tag.NBTPayload.ByteArray)

	case TAG_STRING:
		WriteString(output, tag.NBTPayload.Str)

	case TAG_LIST:
		WriteList(output, tag.NBTPayload.List)

	case TAG_COMPOUND:
		WriteCompound(output, tag.NBTPayload.Compound)

	case TAG_INT_ARRAY:
		WriteIntArray(output, tag.NBTPayload.IntArray)

	case TAG_LONG_ARRAY:
		WriteLongArray(output, tag.NBTPayload.LongArray)
	default:
		panic("UNREACHEABLE")
	}
}
