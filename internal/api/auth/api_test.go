package auth_test

import (
	"context"
	"crypto/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matt-hoiland/architecting/data"
	"github.com/matt-hoiland/architecting/internal/api/auth"
	"github.com/matt-hoiland/architecting/internal/api/auth/mocks"
	itesting "github.com/matt-hoiland/architecting/lib/testing"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewAuthAPI(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	// Mock dependency
	credMock := mocks.NewMockCollector(c)

	// Test Unit
	api := auth.NewAuthAPI(credMock)

	assert.NotNil(t, api)
}

func TestAuthAPI_DeleteCredentials(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	err := api.DeleteCredentials(context.TODO(), nil)
	itesting.AssertMethodImplemented(t, err)
}

func TestAuthAPI_DeleteCredentialsByEmail(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	err := api.DeleteCredentialsByEmail(context.TODO(), "")
	itesting.AssertMethodImplemented(t, err)
}

func TestAuthAPI_DeleteCredentialsByID(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	err := api.DeleteCredentialsByID(context.TODO(), primitive.ObjectID{})
	itesting.AssertMethodImplemented(t, err)
}

func TestAuthAPI_FindCredentials(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	_, err := api.FindCredentials(context.TODO(), "", "")
	itesting.AssertMethodImplemented(t, err)
}

func TestAuthAPI_FindCredentialsByEmail(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	_, err := api.FindCredentialsByEmail(context.TODO(), "")
	itesting.AssertMethodImplemented(t, err)
}

func TestAuthAPI_FindCredentialsByID(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	_, err := api.FindCredentialsByID(context.TODO(), primitive.ObjectID{})
	itesting.AssertMethodImplemented(t, err)
}

func TestAuthAPI_InsertCredentials(t *testing.T) {
	itesting.RedirectLogs()

	type Test struct {
		Name            string
		Contexter       func() context.Context
		Creds           *data.Credentials
		ObjectID        primitive.ObjectID
		Err             error
		SetExpectations func(test Test, coll *mocks.MockCollector)
	}

	var (
		testID    primitive.ObjectID = primitive.NewObjectID()
		testEmail string             = "noone@nowhere.com"
		testHash  []byte             = make([]byte, 32)
		testSalt  []byte             = make([]byte, 32)
	)

	rand.Read(testHash[:])
	rand.Read(testSalt[:])

	tests := []Test{
		{
			Name: "happy",
			Creds: &data.Credentials{
				Email: testEmail,
				Hash:  testHash,
				Salt:  testSalt,
			},
			ObjectID: testID,
			SetExpectations: func(test Test, coll *mocks.MockCollector) {
				coll.EXPECT().
					InsertOne(gomock.Any(), test.Creds).
					Return(&mongo.InsertOneResult{InsertedID: test.ObjectID}, test.Err).
					Times(1)
			},
		},
		{
			Name:     "nil credentials object",
			ObjectID: primitive.NilObjectID,
			Err:      mongo.ErrNilDocument,
			SetExpectations: func(test Test, coll *mocks.MockCollector) {
				coll.EXPECT().
					InsertOne(gomock.Any(), test.Creds).
					Return(nil, test.Err).
					Times(1)
			},
		},
	}
	_ = tests

	for i := range tests {
		test := tests[i]
		t.Run(test.Name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			if test.Contexter == nil {
				test.Contexter = func() context.Context { return context.TODO() }
			}

			credMock := mocks.NewMockCollector(c)
			api := auth.NewAuthAPI(credMock)
			if test.SetExpectations != nil {
				test.SetExpectations(test, credMock)
			}

			id, err := api.InsertCredentials(test.Contexter(), test.Creds)

			itesting.AssertMethodImplemented(t, err)
			assert.Equal(t, test.ObjectID, id)
			assert.Equal(t, test.Err, err)
		})
	}
}

func TestAuthAPI_UpdateCredentials(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	_, err := api.UpdateCredentials(context.TODO(), nil)
	itesting.AssertMethodImplemented(t, err)
}
