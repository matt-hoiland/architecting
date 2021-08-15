package auth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matt-hoiland/architecting/internal/api/auth"
	"github.com/matt-hoiland/architecting/internal/api/auth/mocks"
	"github.com/stretchr/testify/assert"
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
