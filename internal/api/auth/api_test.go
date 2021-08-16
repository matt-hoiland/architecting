package auth_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matt-hoiland/architecting/internal/api/auth"
	"github.com/matt-hoiland/architecting/internal/api/auth/mocks"
	itesting "github.com/matt-hoiland/architecting/lib/testing"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	_, err := api.InsertCredentials(context.TODO(), nil)
	itesting.AssertMethodImplemented(t, err)

}

func TestAuthAPI_UpdateCredentials(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	credMock := mocks.NewMockCollector(c)
	api := auth.NewAuthAPI(credMock)

	_, err := api.UpdateCredentials(context.TODO(), nil)
	itesting.AssertMethodImplemented(t, err)

}
