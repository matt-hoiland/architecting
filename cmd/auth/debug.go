package main

import (
	"crypto/rand"

	"github.com/matt-hoiland/architecting/data"
	"github.com/matt-hoiland/architecting/internal/api/auth"
	log "github.com/sirupsen/logrus"
)

const (
	HashLength = 32
	SaltLength = 32
)

func debug(api *auth.AuthAPI) {
	matt := data.UserCredentials{
		Email: "noone@nowhere.com",
		Hash:  make([]byte, HashLength),
		Salt:  make([]byte, SaltLength),
	}
	rand.Read(matt.Hash)
	rand.Read(matt.Salt)

	id, err := api.InsertCredentials(&matt)
	if err != nil {
		log.Error(err)
	}
	log.Debug("id: " + id.Hex())
	doc, err := api.FindCredentialsByID(id)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("doc: %v", doc)
}
