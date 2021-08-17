package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CredentialsCollection names the credentials collection
	CredentialsCollection = "credentials"

	// TODO: Tie hash and salt minimum lengths to the hashing algorithm's output.
	// HashMinLength is the minimum number of bytes a hash should be.
	HashMinLength = 32
	// SaltMinLength is the minimum number of bytes a salt should be.
	SaltMinLength = 32
)

// AuthCredentials represents the document model stored in mongo.
type AuthCredentials struct {
	// ID is the mongo ObjectID of the document.
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	// Email stores the email address the user uses to login. It must be RFC5322 compliant and unique
	// to the collection. AuthCredentialsSchema enforces FRC5322 compliance. The schema has no way at
	// this time to validate uniqueness. Email can be changed as long as compliance and uniqueness are
	// preserved.
	Email string `bson:"email" json:"email"`
	// Hash stores the salted and hashed password the user uses to login. Its minimum length is set
	// by HashMinLength, which AuthCredentialsSchema enforces. It can be changed at the user's request.
	Hash ByteSlice `bson:"hash" json:"-"`
	// Salt stores the salt used to hash the user's password. Its minimum length is set by SaltMinLength,
	// which AuthCredentialsSchema enforces. A new salt will be generated whenever Hash is changed.
	Salt ByteSlice `bson:"salt" json:"-"`
	// Validated stores whether the user has validated their email address. Its default value is false
	// and will be reset to false if Email is updated.
	Validated bool `bson:"validated" json:"validated"`
	// TimestampToExpire stores the date afterwhich this credentials document will be invalidated. The
	// user must submit a new password to continue using these credentials. This field is to future
	// proof an optional expiration policy.
	TimestampToExpire time.Time `bson:"ts_toexpire,omitempty" json:"ts_toexpire,omitempty"`
	// TimestampCreated stores the time an instance of AuthCredentials is allocated by NewAuthCredentials.
	TimestampCreated time.Time `bson:"ts_created" json:"ts_created"`
	// TimestampUpdated stores the approximate time a change to the document was made last.
	TimestampUpdated time.Time `bson:"ts_updated,omitempty" json:"ts_updated,omitempty"`
}

// NewAuthCredentials builds and returns a new instance of AuthCredentials with the given data. It sets
// the TimestampCreated field on the instance automatically.
func NewAuthCredentials(email string, hash, salt ByteSlice) *AuthCredentials {
	panic("unimplemented")
}

// AuthCredentialsSchema is a jsonSchema document passed to CreateCollection when constructing
// the Credentials collection.
var AuthCredentialsSchema = primitive.M{
	"$jsonSchema": primitive.M{
		"bsonType": "object",
		"required": primitive.A{"email", "hash", "salt", "validated"},
		"properties": primitive.M{
			"_id": primitive.M{
				"bsonType": "objectId",
			},
			"email": primitive.M{
				"bsonType": "string",
				"pattern":  "^(([^<>()\\[\\]\\\\.,;:\\s@\"]+(\\.[^<>()\\[\\]\\\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$",
			},
			"hash": primitive.M{
				"bsonType":  "string",
				"minLength": HashMinLength * 2,
				"pattern":   "[0-9a-fA-F]+",
			},
			"salt": primitive.M{
				"bsonType":  "string",
				"minLength": SaltMinLength * 2,
				"pattern":   "[0-9a-fA-F]+",
			},
			"validated": primitive.M{
				"bsonType": "bool",
			},
			"ts_created": primitive.M{
				"bsonType": "timestamp",
			},
			"ts_updated": primitive.M{
				"bsonType": "timestamp",
			},
		},
	},
}
