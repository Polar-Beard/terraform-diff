package main

import (
	"flag"
	"fmt"
)

var version = "unknown"

func main() {
	var printVersion bool
	flag.BoolVar(&printVersion, "v", false, "prints the version")
	flag.Parse()

	if printVersion {
		fmt.Println(version)
	}
}
