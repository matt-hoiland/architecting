package testing

import (
	"bytes"

	log "github.com/sirupsen/logrus"
)

func RedirectLogs() *bytes.Buffer {
	buf := &bytes.Buffer{}
	log.SetOutput(buf)
	return buf
}
