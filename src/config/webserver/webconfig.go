package config

import (
	"path/filepath"
)

var (
	LOGDIR  string = "/home/asn/projects/goprojects/gofirst/cmd/webserver/logs"
	LOGFILE string = filepath.Join(LOGDIR, "gofirst.log")
)

const (
	SERVER_PORT = 4505
	// SERVER_IP   string = "0.0.0.0"
)
