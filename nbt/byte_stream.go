package nbt

import (
	"bytes"
	"encoding/binary"
	"math"
)

type NBTByteStream struct {
	stream *bytes.Reader
}

func (bs NBTByteStream) ReadLongArray() ([]int64, error) {
	length, err := bs.ReadIntBE()
	if err != nil {
		return nil, err
	}

	arr := make([]int64, 0)
	for i := int32(0); i < length; i++ {
		data, err := bs.ReadLongBE()
		if err != nil {
			return nil, err
		}

		arr = append(arr, data)
	}

	return arr, nil
}

func (bs NBTByteStream) ReadIntArray() ([]int32, error) {
	length, err := bs.ReadIntBE()
	if err != nil {
		return nil, err
	}

	arr := make([]int32, 0)
	for i := int32(0); i < length; i++ {
		data, err := bs.ReadIntBE()
		if err != nil {
			return nil, err
		}

		arr = append(arr, data)
	}

	return arr, nil
}

func (bs NBTByteStream) ReadByteArray() ([]byte, error) {
	length, err := bs.ReadIntBE()
	if err != nil {
		return nil, err
	}

	arr := make([]byte, length)
	for i := int32(0); i < length; i++ {
		data, err := bs.ReadByte()
		if err != nil {
			return nil, err
		}

		arr = append(arr, data)
	}

	return arr, nil
}

func (bs NBTByteStream) ReadFloatBE() (float32, error) {
	f := make([]byte, 4)
	_, err := bs.stream.Read(f)
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(binary.BigEndian.Uint32(f)), nil
}

func (bs NBTByteStream) ReadDoubleBE() (float64, error) {
	d := make([]byte, 8)
	_, err := bs.stream.Read(d)
	if err != nil {
		return 0, err
	}

	return math.Float64frombits(binary.BigEndian.Uint64(d)), nil
}

func (bs NBTByteStream) ReadString() (string, error) {
	strLength, err := bs.ReadShortBE()
	if err != nil {
		return "", err
	}

	name := make([]byte, strLength)
	_, err = bs.stream.Read(name)
	return string(name), err
}

func (bs NBTByteStream) ReadByte() (byte, error) {
	return bs.stream.ReadByte()
}

func (bs NBTByteStream) ReadShortBE() (int16, error) {
	s := make([]byte, 2)
	_, err := bs.stream.Read(s)
	if err != nil {
		return 0, err
	}

	sr := int16(s[0])<<8 | int16(s[1])
	return sr, nil
}

func (bs NBTByteStream) ReadIntBE() (int32, error) {
	i := make([]byte, 4)
	_, err := bs.stream.Read(i)
	if err != nil {
		return 0, err
	}

	ir := int32(i[0])<<24 | int32(i[1])<<16 | int32(i[2])<<8 | int32(i[3])
	return ir, nil
}

func (bs NBTByteStream) ReadLongBE() (int64, error) {
	l := make([]byte, 8)
	_, err := bs.stream.Read(l)
	if err != nil {
		return 0, err
	}

	lr := int64(l[0])<<56 | int64(l[1])<<48 | int64(l[2])<<40 | int64(l[3])<<32 | int64(l[4])<<24 | int64(l[5])<<16 | int64(l[6])<<8 | int64(l[7])
	return lr, nil
}
