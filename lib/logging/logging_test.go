package logging_test

import (
	"bytes"
	"testing"

	"github.com/matt-hoiland/architecting/lib/logging"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLevelFromtString(t *testing.T) {
	invalid := "invalid"
	upinvalid := "INVALID"
	tests := []struct {
		Name          string
		Values        map[string]logrus.Level
		InvalidOption *string
	}{
		{
			Name: "upper case",
			Values: map[string]logrus.Level{
				"PANIC": logrus.PanicLevel,
				"FATAL": logrus.FatalLevel,
				"ERROR": logrus.ErrorLevel,
				"WARN":  logrus.WarnLevel,
				"INFO":  logrus.InfoLevel,
				"DEBUG": logrus.DebugLevel,
				"TRACE": logrus.TraceLevel,
			},
		},
		{
			Name: "lower case",
			Values: map[string]logrus.Level{
				"panic": logrus.PanicLevel,
				"fatal": logrus.FatalLevel,
				"error": logrus.ErrorLevel,
				"warn":  logrus.WarnLevel,
				"info":  logrus.InfoLevel,
				"debug": logrus.DebugLevel,
				"trace": logrus.TraceLevel,
			},
		},
		{
			Name:          "invalid level",
			InvalidOption: &invalid,
		},
		{
			Name:          "invalid level upper case",
			InvalidOption: &upinvalid,
		},
	}

	for i := range tests {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			for in, out := range tc.Values {
				rcv := logging.LevelFromString(in)
				assert.Equal(t, out, rcv)
			}
			if tc.InvalidOption != nil {
				buffer := bytes.Buffer{}
				logrus.SetOutput(&buffer)
				logrus.SetLevel(logrus.WarnLevel)
				rcv := logging.LevelFromString(*tc.InvalidOption)
				assert.Equal(t, logrus.InfoLevel, rcv)
				assert.Contains(t, buffer.String(), "Invalid level provided: "+*tc.InvalidOption)
			}
		})
	}
}
