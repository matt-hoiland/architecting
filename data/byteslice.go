package data

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// ByteSlice wraps the native []byte type for custom bson marshalling to and from bson strings.
type ByteSlice []byte

// MarshalBSONValue implements the bson.ValueMarshaller interface for ByteSlice.
func (bs *ByteSlice) MarshalBSONValue() (bsontype.Type, []byte, error) {
	hexData := hex.EncodeToString(*bs)
	btype, data, err := bson.MarshalValue(hexData)
	return btype, data, err
}

// UnmarshalBSONValue implements the bson.ValueUnmarshaller interface for ByteSlice.
func (bs *ByteSlice) UnmarshalBSONValue(btype bsontype.Type, bdata []byte) error {
	// bson string format: int34(4 bytes, little endian) (*bytes) \x00(null-terminator)
	buf := bytes.NewBuffer(bdata)
	// Read int32 encoded array length
	var n int32
	err := binary.Read(buf, binary.LittleEndian, &n)
	if err != nil {
		return err
	}
	// Read remainging contents of bdata
	data := make([]byte, n)
	_, err = io.ReadFull(buf, data)
	if err != nil {
		return err
	}
	// Convert data to go string (dropping null-terminator) and decode from hex
	*bs, err = hex.DecodeString(string(data[:len(data)-1]))
	if err != nil {
		return err
	}
	// This would be soooo much easier if the driver had a symmetric UnmarshalValue
	// line in the MarshalBSONValue function above; but, alas, bson strings are
	// nullable and go strings are not.
	return nil
}
