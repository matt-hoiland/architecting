package data_test

import (
	"bytes"
	"crypto/rand"
	"testing"
	"time"

	"github.com/matt-hoiland/architecting/data"
	"github.com/stretchr/testify/assert"
)

func TestNewCredentials(t *testing.T) {
	// Test values
	var (
		testTime                 = MustParse("1992-05-28T23:12:11+05:00")
		testEmail                = "matt@hoiland.com"
		testHash  data.ByteSlice = make([]byte, data.HashMinLength)
		testSalt  data.ByteSlice = make([]byte, data.SaltMinLength)
		// episolon is the acceptable variance in Test 2
		episolon = time.Second
	)

	// Initialize hash and salt
	rand.Read(testHash)
	rand.Read(testSalt)

	t.Run("with mocked timestamper", func(t *testing.T) {
		data.Timestamper = func() time.Time { return testTime }
		creds := data.NewCredentials(testEmail, testHash, testSalt)
		if !assert.NotNil(t, creds) {
			t.FailNow()
		}
		assert.Equal(t, testTime, creds.TimestampCreated)
		assert.Equal(t, testEmail, creds.Email)
		assert.True(t, bytes.Equal(testHash, creds.Hash))
		assert.True(t, bytes.Equal(testSalt, creds.Salt))
	})

	t.Run("with now func", func(t *testing.T) {
		data.Timestamper = time.Now
		creds := data.NewCredentials(testEmail, testHash, testSalt)
		if !assert.NotNil(t, creds) {
			t.FailNow()
		}
		assert.True(t, time.Since(creds.TimestampCreated) < episolon)
		assert.Equal(t, testEmail, creds.Email)
		assert.True(t, bytes.Equal(testHash, creds.Hash))
		assert.True(t, bytes.Equal(testSalt, creds.Salt))
	})
}

// MustParse is a helper function which guarantees a correct time.Time value for the given
// timestamp. It panics otherwise.
func MustParse(timestamp string) time.Time {
	time, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		panic(err)
	}
	return time
}
