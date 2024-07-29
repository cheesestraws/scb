package main

import (
	"fmt"
	"os"
	"flag"
)

var args []string

func usage() {
	fmt.Printf("(help message goes here)\n")
}

func main() {
	into := flag.String("into", ".", "put files in this dir")
	flag.Parse()
	
	if *into != "." {
		err := os.Chdir(*into)
		if err != nil {
			fmt.Printf("error: %v\n\n", err)
			return
		}
	}

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
