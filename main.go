///////////////////////////////////////////////////////////////////////////////
//	main.go  -  Dec-5-2022  -  aldebap
//
//	Entry point for Go-DMig
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	versionInfo string = "Go-DMig 0.1"
)

//	main entry point for Go-DMig CLI
func main() {
	var (
		version bool
	)

	//	CLI arguments
	flag.BoolVar(&version, "version", false, "show Go-DMig version")

	flag.Parse()

	//	version option
	if version {
		fmt.Printf("%s\n", versionInfo)
		return
	}

	//	get the Go-DMig configuration file name
	dmigFileName := flag.Arg(0)
	if len(dmigFileName) == 0 {
		fmt.Fprintf(os.Stderr, "[error] missing Go-DMig configuration file name\n")
		os.Exit(-1)
	}
}
