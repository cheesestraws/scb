package main

import (
	"flag"
	"fmt"
	"os"
)

var args []string

func usage() {
	fmt.Printf("%s fetch <kind> <device> <user> <pass> - print configuration to stdout\n", os.Args[0])
	fmt.Printf("%s [-into <dir>] download <kind> <device> <user> <pass> - download a single configuration\n", os.Args[0])
	fmt.Printf("%s [-into <dir>] download-all <tomlfile> - download configs for all devices in tomlfile\n", os.Args[0])
	fmt.Printf("%s toml-check <tomlfile> - check approximate validity of toml file\n", os.Args[0])
	fmt.Printf("\n")
	flag.PrintDefaults()
}

var into *string

func main() {
	flag.Usage = usage
	into = flag.String("into", ".", "put files in this dir")
	flag.Parse()

	var sc string
	var err error
	sc, args, err = extractSubcommand(flag.Args())

	if err != nil {
		usage()
		return
	}

	err = runSubcommand(sc, args)
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		usage()
		return
	}
}
