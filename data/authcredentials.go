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

const (
	HashMinLength = 32
	SaltMinLength = 32
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

var AuthCredentialsSchema = primitive.M{
	"$jsonSchema": primitive.M{
		"bsonType":             "object",
		"required":             primitive.A{"email", "hash", "salt", "validated"},
		"additionalProperties": false,
		"properties": primitive.M{
			"_id": primitive.M{
				"bsonType": "objectId",
			},
			"email": primitive.M{
				"bsonType":    "string",
				"description": "email address of the user; uniquely identifies this document",
				"pattern":     "^(([^<>()\\[\\]\\\\.,;:\\s@\"]+(\\.[^<>()\\[\\]\\\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$",
			},
			"hash": primitive.M{
				"bsonType":    "string",
				"description": "salted and hash user password stored in hexadecimal",
				"minLength":   HashMinLength * 2,
				"pattern":     "[0-9a-fA-F]+",
			},
			"salt": primitive.M{
				"bsonType":    "string",
				"description": "salt used in password hash",
				"minLength":   SaltMinLength * 2,
				"pattern":     "[0-9a-fA-F]+",
			},
			"validated": primitive.M{
				"bsonType":    "bool",
				"description": "whether the email address has been validated",
			},
			"ts_updated": primitive.M{
				"bsonType":    "timestamp",
				"description": "if exists, represents the last time a change was made to the document",
			},
		},
	},
}
