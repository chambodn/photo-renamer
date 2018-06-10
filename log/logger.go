/*
Log package
*/

package log

import (
	"go.uber.org/zap"
)

// Logger is the default info logger
var Logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()

	l, err := config.Build()
	if err != nil {
		panic(err)
	}
	Logger = l
}
