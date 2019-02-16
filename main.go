package main

import (
	"github.com/hirokisan/go-sandbox/lib/logger"
)

func main() {
	logger := logger.Logger(logger.NewLogrusLogger())
	logger.Info("aaaa")
}
