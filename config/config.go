package config

import "github.com/georgekaran/go-jwt-server/util/file"

var ConfigMap map[string]string

func init() {
	ConfigMap = file.ToMap("config.properties")
}