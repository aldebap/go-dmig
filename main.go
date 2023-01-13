///////////////////////////////////////////////////////////////////////////////
//	main.go  -  Dec-5-2022  -  aldebap
//
//	Entry point for Go-DMig
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	migration "github.com/aldebap/go-dmig/migration"
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

	//	open Go-DMig config file and load it
	dmigFile, err := os.Open(dmigFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] opening Go-DMig file: "+err.Error()+"\n")
		os.Exit(-1)
	}
	defer dmigFile.Close()

	//	load the config file
	dmig, err := migration.LoadConfigFile(bufio.NewReader(dmigFile))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] fail loading Go-DMig file: "+err.Error()+"\n")
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stdout, ">>> Migration config: "+dmig.Description+"\n")
}
