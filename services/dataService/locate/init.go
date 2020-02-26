package locate

import (
	"os"

	"tonyxiong.top/gostorage/pkg/logs"
)

// inner package logger
var logger *logs.Logger

func init() {
	// initial logger
	logger = logs.NewLogger(os.Stdout, map[string]string{
		"appName": "dataService",
		"package": "locate",
	})
}
