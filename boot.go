package main

import "os"

func Boot() {
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "--help" {
			printHelp()
			return
		}
		handleCLI(args)
		return
	}
	interactiveMode()
}
