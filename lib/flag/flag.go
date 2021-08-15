package flag

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Parse calls stdlib's flag.Parse function.
func Parse() {
	flag.Parse()
}

// String sets up a command line flag/environment variable combination.
func String(name, envKey, defaultValue, usage string) *string {
	return flag.String(name, LookupEnvOrString(envKey, defaultValue), appendDefault(usage, defaultValue))
}

// Int sets up a command line flag/environment variable combination.
func Int(name, envKey string, defaultValue int, usage string) *int {
	return flag.Int(name, LookupEnvOrInt(envKey, defaultValue), appendDefault(usage, defaultValue))
}

func appendDefault(usage string, defaultValue interface{}) string {
	return fmt.Sprintf("%s [default: %v]", usage, defaultValue)
}

// LookupEnvOrString checks the environment for the value of the given key or
// passes through the default value provided.
func LookupEnvOrString(envKey, defaultValue string) string {
	if val, ok := os.LookupEnv(envKey); ok {
		return val
	}
	return defaultValue
}

// LookupEnvOrInt checks the environment for the value of the given key and
// parses it to int or passes through the default value provided.
func LookupEnvOrInt(envKey string, defaultValue int) int {
	val, ok := os.LookupEnv(envKey)
	if !ok {
		return defaultValue
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		log.WithFields(log.Fields{
			"key":     envKey,
			"val":     val,
			"default": defaultValue,
		}).Warn("Non-int value given for an int envar")
		return defaultValue
	}
	return i
}
