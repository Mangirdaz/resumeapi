package main

import (
	"os"
	"strconv"
)

type KeyValueStorageConfig struct {
	ip             string
	port           string
	timeoutseconds int
}

func InitKeyValueStorageConfig() KeyValueStorageConfig {

	if os.Getenv("KEYVAL_STORAGE_IP") == "" {
		config.ip = "172.17.0.2"
	} else {
		config.ip = os.Getenv("KEYVAL_STORAGE_PORT")
	}

	if os.Getenv("KEYVAL_STORAGE_PORT") == "" {
		config.port = "8500"
	} else {
		config.port = os.Getenv("KEYVAL_STORAGE_PORT")
	}

	if os.Getenv("KEYVAL_STORAGE_TIMEOUT") == "" {
		config.timeoutseconds = 10
	} else {
		config.timeoutseconds, _ = strconv.Atoi(os.Getenv("KEYVAL_STORAGE_TIMEOUT"))
	}

	return config
}
