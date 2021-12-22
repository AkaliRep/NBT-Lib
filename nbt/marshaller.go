package nbt

import (
	"bytes"
	"encoding/binary"
	"log"
	"math"
)

// func Marshal(t interface{}) []byte {
// }

// func Unmarshal(t interface{}) []byte {
// }

func ReadFile(filename string) NBTTag {
	file, err := DecodeFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(file)
	return Read(r)
}

func Read(file *bytes.Reader) NBTTag {
	return readImplicitCompound(file)
}

func readImplicitCompound(file *bytes.Reader) NBTTag {
	tagType, err := file.ReadByte()
	if err != nil {
		log.Fatalf("Couldn't read tagType, %s\n", err)
	}

	if tagType != TAG_COMPOUND {
		log.Fatalf("Expected TAG_COMPOUND but got %s\n", TagTypeToString(tagType))
	}

	a := readName(file)

	return NBTTag{
		NBTType: TAG_COMPOUND,
		Name:    a,
		NBTPayload: NBTPayloadUnion{
			Compound: getNbtCompoundBody(file),
		},
	}
}

func getNbtCompoundBody(file *bytes.Reader) map[string]NBTTag {
	nbtTagCompound := make(map[string]NBTTag)
	nextTag, err := file.ReadByte()
	if err != nil {
		log.Fatal("Error reading the name tag, ", err)
	}

	for nextTag != TAG_END {
		name := readName(file)
		nbtTagCompound[name] = readNbtTag(nextTag, file)

		nextTag, err = file.ReadByte()
		if err != nil {
			log.Fatal("Couldn't read next tag, ", err)
		}
	}

	return nbtTagCompound
}

func readNbtTag(tag byte, file *bytes.Reader) NBTTag {
	switch tag {
	case TAG_BYTE:
		return readByte(file)
	case TAG_SHORT:
		return readShort(file)
	case TAG_INT:
		return readInt(file)
	case TAG_LONG:
		return readLong(file)
	case TAG_FLOAT:
		return readFloat(file)
	case TAG_DOUBLE:
		return readDouble(file)
	case TAG_BYTE_ARRAY:
		return readByteArray(file)
	case TAG_STRING:
		return readString(file)
	case TAG_LIST:
		return readList(file)
	case TAG_COMPOUND:
		return readCompound(file)
	case TAG_INT_ARRAY:
		return readIntArray(file)
	case TAG_LONG_ARRAY:
		return readLongArray(file)
	default:
		panic("UNREACHEABLE")
	}
}

func readByte(file *bytes.Reader) NBTTag {
	return NBTTag{
		NBTType: TAG_BYTE,
		NBTPayload: NBTPayloadUnion{
			S8: readByteHelper(file),
		},
	}
}

func readShort(file *bytes.Reader) NBTTag {
	return NBTTag{
		NBTType: TAG_SHORT,
		NBTPayload: NBTPayloadUnion{
			S16: readShortHelper(file),
		},
	}
}

func readInt(file *bytes.Reader) NBTTag {
	return NBTTag{
		NBTType: TAG_INT,
		NBTPayload: NBTPayloadUnion{
			S32: readIntHelper(file),
		},
	}
}

func readLong(file *bytes.Reader) NBTTag {
	return NBTTag{
		NBTType: TAG_LONG,
		NBTPayload: NBTPayloadUnion{
			S64: readLongHelper(file),
		},
	}
}

func readFloat(file *bytes.Reader) NBTTag {
	f := make([]byte, 4)
	_, err := file.Read(f)
	if err != nil {
		log.Fatal("Error reading float, ", err)
	}

	fr := math.Float32frombits(binary.BigEndian.Uint32(f))
	return NBTTag{
		NBTType: TAG_FLOAT,
		NBTPayload: NBTPayloadUnion{
			F32: fr,
		},
	}
}

func readDouble(file *bytes.Reader) NBTTag {
	d := make([]byte, 8)
	_, err := file.Read(d)
	if err != nil {
		log.Fatal("Error reading double, ", err)
	}

	dr := math.Float64frombits(binary.BigEndian.Uint64(d))
	return NBTTag{
		NBTType: TAG_DOUBLE,
		NBTPayload: NBTPayloadUnion{
			F64: dr,
		},
	}
}

func readByteArray(file *bytes.Reader) NBTTag {
	length := int(readIntHelper(file))
	arr := make([]byte, length)
	for i := 0; i < length; i++ {
		arr = append(arr, readByteHelper(file))
	}

	return NBTTag{
		NBTType: TAG_INT_ARRAY,
		NBTPayload: NBTPayloadUnion{
			ByteArray: arr,
		},
	}
}

func readString(file *bytes.Reader) NBTTag {
	return NBTTag{
		NBTType: TAG_STRING,
		NBTPayload: NBTPayloadUnion{
			Str: readName(file),
		},
	}
}

func readList(file *bytes.Reader) NBTTag {
	panic("UNIMPLEMENTED")
}

func readCompound(file *bytes.Reader) NBTTag {
	return NBTTag{
		NBTType: TAG_COMPOUND,
		NBTPayload: NBTPayloadUnion{
			Compound: getNbtCompoundBody(file),
		},
	}
}

func readIntArray(file *bytes.Reader) NBTTag {
	length := int(readIntHelper(file))
	arr := make([]int32, length)
	for i := 0; i < length; i++ {
		arr = append(arr, readIntHelper(file))
	}

	return NBTTag{
		NBTType: TAG_INT_ARRAY,
		NBTPayload: NBTPayloadUnion{
			IntArray: arr,
		},
	}
}

func readLongArray(file *bytes.Reader) NBTTag {
	length := int(readIntHelper(file))
	arr := make([]int64, length)
	for i := 0; i < length; i++ {
		arr = append(arr, readLongHelper(file))
	}

	return NBTTag{
		NBTType: TAG_INT_ARRAY,
		NBTPayload: NBTPayloadUnion{
			LongArray: arr,
		},
	}
}

func readName(file *bytes.Reader) string {
	nameLength := make([]byte, 2)
	_, err := file.Read(nameLength)
	if err != nil {
		log.Fatal("Couldn't read file", err)
	}

	nameLengthAsUint := binary.BigEndian.Uint16(nameLength)
	name := make([]byte, nameLengthAsUint)
	_, err = file.Read(name)
	if err != nil {
		log.Fatal("Couldn't read name, ", err)
	}

	return string(name)
}

func readByteHelper(file *bytes.Reader) byte {
	b, err := file.ReadByte()
	if err != nil {
		log.Fatal("Error reading byte, ", err)
	}

	return b
}

func readShortHelper(file *bytes.Reader) int16 {
	s := make([]byte, 2)
	_, err := file.Read(s)
	if err != nil {
		log.Fatal("Error reading short, ", err)
	}

	sr := int16(s[0])<<8 | int16(s[1])
	return sr
}

func readIntHelper(file *bytes.Reader) int32 {
	i := make([]byte, 4)
	_, err := file.Read(i)
	if err != nil {
		log.Fatal("Error reading int, ", err)
	}

	ir := int32(i[0])<<24 | int32(i[1])<<16 | int32(i[2])<<8 | int32(i[3])
	return ir
}

func readLongHelper(file *bytes.Reader) int64 {
	l := make([]byte, 8)
	_, err := file.Read(l)
	if err != nil {
		log.Fatal("Error reading long, ", err)
	}

	lr := int64(l[0])<<56 | int64(l[1])<<48 | int64(l[2])<<40 | int64(l[3])<<32 | int64(l[4])<<24 | int64(l[5])<<16 | int64(l[6])<<8 | int64(l[7])
	return lr
}
