package main

import (
	"os"
)

type KeyValueStorageConfig struct {
	Ip   string
	Port string
}

type ErrorMessage struct {
	Code    int
	Message string
}

func InitKeyValueStorageConfig() KeyValueStorageConfig {

	var config KeyValueStorageConfig

	if os.Getenv("KEYVAL_STORAGE_IP") == "" {
		config.Ip = "172.17.0.2"
	} else {
		config.Ip = os.Getenv("KEYVAL_STORAGE_PORT")
	}

	if os.Getenv("KEYVAL_STORAGE_PORT") == "" {
		config.Port = "8500"
	} else {
		config.Port = os.Getenv("KEYVAL_STORAGE_PORT")
	}

	return config
}
