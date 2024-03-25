package config

import "os"

var DbServer = "localhost"
var DbPort = 1433
var DbUser = "sa"
var DbPassword = "leehi357!@"
var DbName = "master"

var PORT = ":9090"

var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"

func InitAppConfig() {
	staticRouteEnv := os.Getenv("STATIC_ROUTE")
	if staticRouteEnv != "" {
		STATIC_ROUTE = staticRouteEnv
	}

	staticDirEnv := os.Getenv("STATIC_DIR")
	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}
}
