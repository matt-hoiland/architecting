package main

import (
	"bytes"
	"context"
	"crypto/rand"

	"github.com/matt-hoiland/architecting/data"
	"github.com/matt-hoiland/architecting/internal/api/auth"
	log "github.com/sirupsen/logrus"
)

const (
	HashLength = 32
	SaltLength = 32
)

func debug(ctx context.Context, api *auth.AuthAPI) {
	matt := data.AuthCredentials{
		Email: "noone@nowhere.com",
		Hash:  make([]byte, HashLength),
		Salt:  make([]byte, SaltLength),
	}
	rand.Read(matt.Hash)
	rand.Read(matt.Salt)

	id, err := api.InsertCredentials(ctx, &matt)
	if err != nil {
		log.Error(err)
	}
	log.Debug("id: " + id.Hex())
	doc, err := api.FindCredentialsByID(ctx, id)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("crd: %v", matt)
	log.Debugf("doc: %v", doc)
	log.Debugf("hash equals: %t, salt equals: %t", bytes.Equal(matt.Hash, doc.Hash), bytes.Equal(matt.Salt, doc.Salt))
}
