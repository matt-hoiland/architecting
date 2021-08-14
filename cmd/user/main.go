package main

import (
	"github.com/matt-hoiland/architecting/lib/flag"
	"github.com/matt-hoiland/architecting/lib/logging"
	log "github.com/sirupsen/logrus"
)

const (
	ServiceName     = "UserService"
	DefaultLogLevel = "INFO"
)

var (
	logLevel = flag.String("log-level", "LOG_LEVEL", DefaultLogLevel, "logging level for service")
)

func main() {
	flag.Parse()
	log.SetLevel(logging.LevelFromString(*logLevel))

	log.Info(ServiceName + " starting ...")
	defer log.Info(ServiceName + " closing ...")

	log.Debug("HI! I'm the " + ServiceName)
}
