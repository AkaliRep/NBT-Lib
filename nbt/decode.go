package nbt

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

func DecodeFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return decompressFile(data)
}

const (
	GZIP = 0x1f
	ZLIB = 0x78
)

func decompressFile(file []byte) ([]byte, error) {
	compressType := file[0]
	switch compressType {
	case GZIP:
		return decompressGzip(file)
	case ZLIB:
		return decompressZlib(file)
	default:
		return file, nil
	}
}
func decompressZlib(file []byte) ([]byte, error) {
	panic("UNIMPLEMENTED")
}

func decompressGzip(file []byte) ([]byte, error) {
	gzipr, err := gzip.NewReader(bytes.NewBuffer(file))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(gzipr)
	if err != nil {
		return nil, err
	}
	gzipr.Close()

	return data, nil
}
