package data

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthCredentials represents the document model stored in mongo.
type AuthCredentials struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"        json:"_id,omitempty"`
	Email                string             `bson:"email"                json:"email"`
	Hash                 ByteSlice          `bson:"hash"                 json:"-"`
	Salt                 ByteSlice          `bson:"salt"                 json:"-"`
	Validated            bool               `bson:"validated"            json:"validated"`
	TimestampUpdatedLast time.Time          `bson:"ts_updated,omitempty" json:"ts_updated,omitempty"`
}

// ByteSlice wraps the native byte slice and implements bson.ValueMarshaller and bson.ValueUnmarshaller
type ByteSlice []byte

func (bs *ByteSlice) MarshalBSONValue() (bsontype.Type, []byte, error) {
	hexData := hex.EncodeToString(*bs)
	btype, data, err := bson.MarshalValue(hexData)
	return btype, data, err
}

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
