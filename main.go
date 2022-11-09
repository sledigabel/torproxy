package main

import (
	"fmt"
	"os"

	"github.com/sledigabel/torproxy/pkg/logging"
)

func main() {
	logger, err := logging.GetLogger()
	if err != nil {
		fmt.Printf("Couldn't initialise logging: %s\n", err)
		os.Exit(1)
	}
	logger.Info("hello world")
}
