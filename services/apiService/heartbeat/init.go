package heartbeat

import "tonyxiong.top/gostorage/pkg/logs"

// initial logger
var logger *logs.Logger

func init() {
	logger = logs.NewLogger()
}
