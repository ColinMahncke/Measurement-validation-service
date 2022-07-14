package main

import (
	"os"
)

//fallback for the environment variables
func Getenv(name string, fallback string) string {
	variable, found := os.LookupEnv(name)

	if !found {
		variable = fallback
	}
	return variable
}
