package env

import (
	"fmt"
	"os"
	"strconv"
)

func Require(key string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		panic(fmt.Sprintf("Env variable %s is missing", key))
	}

	return value
}

func Get(key string, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func RequireInt(key string) int {
	intValue, err := strconv.Atoi(Require(key))

	if err != nil {
		panic(fmt.Sprintf("Env variable %s is not of type int", key))
	}

	return intValue
}

func GetInt(key string, defaultValue int) int {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return intValue
}

func RequireBool(key string) bool {
	value := Require(key)

	if value == "true" {
		return true
	}

	if value == "false" {
		return false
	}

	panic(fmt.Sprintf("Env variable %s is not of type bool", key))
}

func GetBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	return value == "true"
}
