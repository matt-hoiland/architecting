package flag

import (
	"flag"
	"os"
)

// Parse calls stdlib's flag.Parse function.
func Parse() {
	flag.Parse()
}

// String setups a command line flag/environment variable combination.
func String(name, envKey, defaultValue, usage string) *string {
	return flag.String(name, LookupEnvOrString(envKey, defaultValue), usage)
}

// LookupEnvOrString checks the environment for the value of the given key or
// passes through the default value provided.
func LookupEnvOrString(envKey, defaultValue string) string {
	if val, ok := os.LookupEnv(envKey); ok {
		return val
	}
	return defaultValue
}
