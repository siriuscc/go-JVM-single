package main

import (
	"fmt"
	"logger"
)

func main() {

	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		logger.DEBUG = cmd.XDebug
		createJVM(cmd).start()
	}
}


